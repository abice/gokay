package gkgen

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

const (
	errorFormat = `errors%s = append(errors%s, %s)`
)

// AddError is a helper method for templates to add an error to a field.
func addError(field string, eString string) (ret string, err error) {
	ret = fmt.Sprintf(errorFormat, field, field, eString)
	return
}

// addFieldError is a helper method for templates to add an error to a field.
func addFieldError(v Validation, eString string) (ret string, err error) {
	name := v.FieldName
	ret = fmt.Sprintf(errorFormat, name, name, eString)
	return
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// IsPtr is a helper method for templates to use to determine if a field is a pointer.
func isPtr(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	// fmt.Printf("ast.Field: %#v\n", field.Type)
	_, ret = field.Type.(*ast.StarExpr)
	return
}

// IsNullable is a helper method for templates to use to determine if a field is a nullable
func isNullable(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	fType := field.Type
	if _, ok := fType.(*ast.Ident); ok {
		fType = getIdentType(fType.(*ast.Ident))
	}

	// fmt.Printf("IsNullable: %s=>%#v\n", f.FieldType, fType)
	switch fType.(type) {
	case *ast.StarExpr:
		ret = true
	case *ast.MapType:
		ret = true
	case *ast.ArrayType:
		ret = true
	case *ast.ChanType:
		ret = true
	case *ast.FuncType:
		ret = true
	case *ast.InterfaceType:
		ret = true
	}
	return
}

// isArray is a helper method for templates to determine if the field is an array
func isArray(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	switch field.Type.(type) {
	case *ast.ArrayType:
		ret = true
	}
	return
}

// isMap is a helper method for templates to determine if the field is a map
func isMap(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	switch field.Type.(type) {
	case *ast.MapType:
		ret = true
	}
	return
}

// GenerationError is a helper method for templates to generate an error if something
// is unsupported.
func GenerationError(s string) (ret interface{}, err error) {
	err = errors.New(s)
	return
}

// isStruct is a helper method for templates to determine if the field is a struct (not a pointer to a struct)
func tmplIsStruct(f Validation) (ret bool, err error) {
	ret = isStruct(f.F)
	return
}

// isStruct is a helper method for templates to determine if the field is a struct (not a pointer to a struct)
func isStruct(field *ast.Field) (ret bool) {
	ret = false
	fType := field.Type
	if _, ok := fType.(*ast.Ident); ok {
		fType = getIdentType(fType.(*ast.Ident))
	}

	switch fType.(type) {
	case *ast.StructType:
		ret = true
	}
	return
}

// tmplIsStructPtr is the template helper method to determine if a field is a *Struct
func tmplIsStructPtr(f Validation) (ret bool, err error) {
	ret = isStructPtr(f.F)
	return
}

// isStructPtr is a helper method for templates to determine if the field is a pointer to a struct
func isStructPtr(field *ast.Field) (ret bool) {
	ret = false
	fType := field.Type
	if star, ok := fType.(*ast.StarExpr); ok {
		fType = getIdentType(star.X.(*ast.Ident))
	}

	switch fType.(type) {
	case *ast.StructType:
		ret = true
	}
	return
}

func getIdentType(ident *ast.Ident) ast.Expr {
	if ident.Obj != nil {
		if spec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
			return spec.Type
		}
	}
	return ident
}
