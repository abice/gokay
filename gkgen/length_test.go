package gkgen

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// LengthTestSuite
type LengthTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *LengthTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestLengthTestSuite(t *testing.T) {
	suite.Run(t, new(LengthTestSuite))
}

var LengthErrorCases = []struct {
	inType      string
	errorString string
}{
	{"int", `Length is not valid on field 'TestField int'`},
	{"int8", `Length is not valid on field 'TestField int8'`},
	{"int16", `Length is not valid on field 'TestField int16'`},
	{"int32", `Length is not valid on field 'TestField int32'`},
	{"int64", `Length is not valid on field 'TestField int64'`},
	{"uint", `Length is not valid on field 'TestField uint'`},
	{"uint8", `Length is not valid on field 'TestField uint8'`},
	{"uint16", `Length is not valid on field 'TestField uint16'`},
	{"uint32", `Length is not valid on field 'TestField uint32'`},
	{"uint64", `Length is not valid on field 'TestField uint64'`},
	{"bool", `Length is not valid on field 'TestField bool'`},
	{"float", `Length is not valid on field 'TestField float'`},
	{"float32", `Length is not valid on field 'TestField float32'`},
	{"float64", `Length is not valid on field 'TestField float64'`},
	{"complex64", `Length is not valid on field 'TestField complex64'`},
	{"complex128", `Length is not valid on field 'TestField complex128'`},
}

func (s *LengthTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range LengthErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"Length=12\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error for '%s' missing", testCase.inType)
		if err != nil {
			s.Contains(err.Error(), testCase.errorString)
		}
	}
}

var LengthSuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"*string", `gokay.LengthString(12, s.TestField); err != nil`},
	{"string", `gokay.LengthString(12, &s.TestField); err != nil`},
	{"[]string", `gokay.LengthString(12, &singleTestField); err != nil`},
	{"[]*string", `gokay.LengthString(12, singleTestField); err != nil`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *LengthTestSuite) TestSuccessCases() {

	format := `package test
	// SomeInterface
	type SomeInterface interface{

	}
	// SomeOtherStruct
	type SomeOtherStruct struct{

	}
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range LengthSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"Length=12\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}

// TestOldFormatSuccessCases will cycle through the test cases for successful calls
// to the Length template and validate that the correct validation has been produced.
func (s *LengthTestSuite) TestOldFormatSuccessCases() {

	format := `package test
	// SomeInterface
	type SomeInterface interface{

	}
	// SomeOtherStruct
	type SomeOtherStruct struct{

	}
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range LengthSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"Length=(12)\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
