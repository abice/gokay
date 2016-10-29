package gkgen

import (
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
	switch field.Type.(type) {
	case *ast.StarExpr:
		ret = true
	case *ast.MapType:
		ret = true
	case *ast.ArrayType:
		ret = true
	case *ast.InterfaceType:
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
