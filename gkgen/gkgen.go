package gkgen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

const (
	validateTag = `valid`
	errorFormat = `errors%s = append(errors%s, %s)`
)

type Generator struct {
	t              *template.Template
	knownTemplates map[string]*template.Template
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator() *Generator {
	g := &Generator{
		knownTemplates: make(map[string]*template.Template),
		t:              template.New("gkgen"),
	}
	g.t.Funcs(map[string]interface{}{
		"CallTemplate": g.CallTemplate,
		"IsPtr":        IsPtr,
		"AddError":     AddFieldError,
	})

	for _, assets := range AssetNames() {
		g.t = template.Must(g.t.Parse(string(MustAsset(assets))))
	}

	g.updateTemplates()

	return g
}

// AddError is a helper method for templates to add an error to a field.
func AddError(field string, eString string) (ret string, err error) {
	ret = fmt.Sprintf(errorFormat, field, field, eString)
	return
}

// AddFieldError is a helper method for templates to add an error to a field.
func AddFieldError(field *ast.Field, eString string) (ret string, err error) {
	name := field.Names[0]
	ret = fmt.Sprintf(errorFormat, name, name, eString)
	return
}

// IsPtr is a helper method for templates to use to determine if a field is a pointer.
func IsPtr(data interface{}) (ret bool, err error) {
	ret = false
	if field, ok := data.(*ast.Field); ok {
		_, ret = field.Type.(*ast.StarExpr)
	} else {
		err = fmt.Errorf("Cannot cast the data past in as an *ast.Field")
	}
	return
}

// CallTemplate is a helper method for the template to call a parsed template but with
// a dynamic name.
func (g *Generator) CallTemplate(name string, data interface{}) (ret string, err error) {
	found := false
	for _, temp := range g.t.Templates() {
		if name == temp.Name() {
			found = true
			break
		}
	}
	buf := bytes.NewBuffer([]byte{})
	if !found {
		fmt.Printf("No template named for '%s' found, ignoring...\n", name)
	} else {
		err = g.t.ExecuteTemplate(buf, name, data)
	}
	ret = buf.String()
	return
}

// GenerateFromFile is responsible for orchestrating the Code generation.  It results in a byte array
// that can be written to any file desired.  It has already had goimports run on the code before being returned.
func (g *Generator) GenerateFromFile(inputFile string) ([]byte, error) {
	f, err := g.parseFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("generate: error parsing input file '%s': %s", inputFile, err)
	}

	structs := g.inspect(f)
	if len(structs) <= 0 {
		return nil, nil
	}

	pkg := f.Name.Name

	vBuff := bytes.NewBuffer([]byte{})
	g.t.ExecuteTemplate(vBuff, "header", map[string]interface{}{"package": pkg})

	for name, st := range structs {
		rules := make(map[string][]string)

		for _, field := range st.Fields.List {
			if field.Tag != nil {
				if strings.Contains(field.Tag.Value, validateTag) {
					str := strings.Trim(field.Tag.Value, "`")
					vals := strings.Split(str, " ")
					for _, val := range vals {
						if strings.HasPrefix(val, validateTag) {
							ruleStr := val[len(validateTag)+2 : len(val)-1]
							fieldRules := strings.Split(ruleStr, ",")
							for _, rule := range fieldRules {
								if _, ok := g.knownTemplates[rule]; ok {
									rules[field.Names[0].Name] = fieldRules
								} else {
									fmt.Printf("Skipping unknown validation template: '%s'", rule)
								}
							}
						}
					}
				}
			}
		}

		data := map[string]interface{}{
			"st":    st,
			"name":  name,
			"rules": rules,
		}

		err := g.t.ExecuteTemplate(vBuff, "struct", data)

		if err != nil {
			if te, ok := err.(template.ExecError); ok {
				return nil, fmt.Errorf("generate: error executing template: %s", te.Err)
			}

			return nil, fmt.Errorf("generate: error executing template: %s", err)
		}
	}

	formatted, err := imports.Process(pkg, vBuff.Bytes(), nil)
	if err != nil {
		fmt.Printf("Error formatting code %s\n\n%s\n", err, string(vBuff.Bytes()))
		err = fmt.Errorf("generate: error formatting code: %s", err)
	}
	return formatted, err

}

func (g *Generator) addTemplate(t string) (err error) {
	g.t = template.Must(g.t.Parse(t))
	g.updateTemplates()
	return
}

func (g *Generator) addTemplateFiles(filenames ...string) (err error) {
	g.t = template.Must(g.t.ParseFiles(filenames...))
	g.updateTemplates()
	return
}

func (g *Generator) updateTemplates() {
	for _, template := range g.t.Templates() {
		g.knownTemplates[template.Name()] = template
	}
}

// parseFile simply calls the go/parser ParseFile function with an empty token.FileSet
func (g *Generator) parseFile(fileName string) (*ast.File, error) {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file given in arguments
	return parser.ParseFile(fset, fileName, nil, parser.ParseComments)
}

func (g *Generator) inspect(f *ast.File) map[string]*ast.StructType {
	structs := make(map[string]*ast.StructType)
	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			if x.Obj != nil {
				// Make sure it's a Type Identifier
				if x.Obj.Kind == ast.Typ {
					// Make sure it's a spec (Type Identifiers can be throughout the code)
					if ts, ok := x.Obj.Decl.(*ast.TypeSpec); ok {
						// Only stsore the struct types (we don't do anything for interfaces)
						if sts, store := ts.Type.(*ast.StructType); store {
							// for _, field := range sts.Fields.List {
							// if field.Type
							// fmt.Printf("Field: %#v\n", field.Type)
							// }
							structs[x.Name] = sts
						}
					}
				}
			}
		}
		// Return true to continue through the tree
		return true
	})

	return structs
}
