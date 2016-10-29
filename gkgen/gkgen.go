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
	validateTag = `valid:`
)

// Validation is a holder for a validation rule within the generation templates
// It actually has the same information as the Field struct, simply for ease of
// access from within the templates.
type Validation struct {
	Name      string
	Param     string
	FieldName string
	F         *ast.Field
}

// Field is used for storing field information.  It holds a reference to the
// original AST field information to help out if needed.
type Field struct {
	Name  string
	F     *ast.Field
	Rules []Validation
	Type  string // Working on getting this figured out
}

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	t              *template.Template
	knownTemplates map[string]*template.Template
	fileSet        *token.FileSet
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator() *Generator {
	g := &Generator{
		knownTemplates: make(map[string]*template.Template),
		t:              template.New("gkgen"),
		fileSet:        token.NewFileSet(),
	}
	g.t.Funcs(map[string]interface{}{
		"CallTemplate": g.CallTemplate,
		"IsPtr":        isPtr,
		"AddError":     addFieldError,
		"IsNullable":   isNullable,
		"typeof":       typeof,
		"isMap":        isMap,
	})

	for _, assets := range AssetNames() {
		g.t = template.Must(g.t.Parse(string(MustAsset(assets))))
	}

	g.updateTemplates()

	return g
}

// CallTemplate is a helper method for the template to call a parsed template but with
// a dynamic name.
func (g *Generator) CallTemplate(rule Validation, data interface{}) (ret string, err error) {
	found := false
	for _, temp := range g.t.Templates() {
		if rule.Name == temp.Name() {
			found = true
			break
		}
	}
	buf := bytes.NewBuffer([]byte{})
	if !found {
		fmt.Printf("No template named for '%s' found, ignoring...\n", rule.Name)
	} else {
		err = g.t.ExecuteTemplate(buf, rule.Name, data)
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
		rules := make(map[string]Field)

		// Go through the fields in the struct and find all the validated tags
		for _, field := range st.Fields.List {
			if field.Tag != nil {
				if strings.Contains(field.Tag.Value, validateTag) {
					// We have a validation flag, make a field
					f := Field{
						F:    field,
						Name: field.Names[0].Name,
					}

					// The AST keeps the rune marker on the string, so we trim them off
					str := strings.Trim(field.Tag.Value, "`")
					// Separate tag types are separated by spaces, so split on that
					vals := strings.Split(str, " ")
					for _, val := range vals {
						// Only parse out the valid: tag
						if strings.HasPrefix(val, validateTag) {
							// Strip off the valid: prefix and the quotation marks
							ruleStr := val[len(validateTag)+1 : len(val)-1]
							// Split on commas for multiple validations
							fieldRules := strings.Split(ruleStr, ",")
							for _, rule := range fieldRules {
								// Rules are able to have parameters,
								// but will have an = in them if that is the case.
								v := Validation{
									Name:      rule,
									FieldName: f.Name,
									F:         f.F,
								}
								if strings.Contains(rule, `=`) {
									// There is a parameter, so get the rule name, and the parameter
									temp := strings.Split(rule, `=`)
									v.Name = temp[0]
									v.Param = temp[1]
								}

								// Only keep the rule if it is a known template
								if _, ok := g.knownTemplates[v.Name]; ok {
									f.Rules = append(f.Rules, v)
								} else {
									fmt.Printf("Skipping unknown validation template: '%s'\n", v.Name)
								}
							}
						}
					}

					// If we have any rules for the field, add it to the map
					if len(f.Rules) > 0 {
						rules[f.Name] = f
					}
				}
			}
		}

		data := map[string]interface{}{
			"st":    st,
			"name":  name,
			"rules": rules,
		}

		err = g.t.ExecuteTemplate(vBuff, "struct", data)

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

// addTemplate is for parsing and add that template string into the template engine.
func (g *Generator) addTemplate(t string) (err error) {
	g.t = template.Must(g.t.Parse(t))
	g.updateTemplates()
	return
}

// AddTemplateFiles will be used during generation when the command line accepts
// user templates to add to the generation.
func (g *Generator) AddTemplateFiles(filenames ...string) (err error) {
	g.t, err = g.t.ParseFiles(filenames...)
	if err == nil {
		g.updateTemplates()
	}
	return
}

// updateTemplates will update the lookup map for validation checks that are
// allowed within the template engine.
func (g *Generator) updateTemplates() {
	for _, template := range g.t.Templates() {
		g.knownTemplates[template.Name()] = template
	}
}

// parseFile simply calls the go/parser ParseFile function with an empty token.FileSet
func (g *Generator) parseFile(fileName string) (*ast.File, error) {
	// Parse the file given in arguments
	return parser.ParseFile(g.fileSet, fileName, nil, parser.ParseComments)
}

// inspect will walk the ast and fill a map of names and their struct information
// for use in the generation template.
func (g *Generator) inspect(f *ast.File) map[string]*ast.StructType {

	structs := make(map[string]*ast.StructType)
	// Inspect the AST and find all structs.
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
