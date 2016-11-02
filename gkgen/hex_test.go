package gkgen

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// HexTestSuite
type HexTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *HexTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestHexTestSuite(t *testing.T) {
	suite.Run(t, new(HexTestSuite))
}

var HexErrorCases = []struct {
	inType      string
	errorString string
}{
	{"int", `Hex is not valid on field 'TestField int'`},
	{"int8", `Hex is not valid on field 'TestField int8'`},
	{"int16", `Hex is not valid on field 'TestField int16'`},
	{"int32", `Hex is not valid on field 'TestField int32'`},
	{"int64", `Hex is not valid on field 'TestField int64'`},
	{"uint", `Hex is not valid on field 'TestField uint'`},
	{"uint8", `Hex is not valid on field 'TestField uint8'`},
	{"uint16", `Hex is not valid on field 'TestField uint16'`},
	{"uint32", `Hex is not valid on field 'TestField uint32'`},
	{"uint64", `Hex is not valid on field 'TestField uint64'`},
	{"bool", `Hex is not valid on field 'TestField bool'`},
	{"float", `Hex is not valid on field 'TestField float'`},
	{"float32", `Hex is not valid on field 'TestField float32'`},
	{"float64", `Hex is not valid on field 'TestField float64'`},
	{"complex64", `Hex is not valid on field 'TestField complex64'`},
	{"complex128", `Hex is not valid on field 'TestField complex128'`},
	{"SomeOtherStruct", `Hex is not valid on field 'TestField SomeOtherStruct'`},
	{"*SomeOtherStruct", `Hex is not valid on field 'TestField *SomeOtherStruct'`},
	{"SomeInterface", `Hex is not valid on field 'TestField SomeInterface'`},
}

func (s *HexTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range HexErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"Hex\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error for '%s' missing", testCase.inType)
		if err != nil {
			s.Contains(err.Error(), testCase.errorString)
		}
	}
}

var HexSuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"*string", `gokay.IsHex(s.TestField); err != nil`},
	{"string", `gokay.IsHex(&s.TestField); err != nil`},
	{"[]string", `gokay.IsHex(&singleTestField); err != nil`},
	{"[]*string", `gokay.IsHex(singleTestField); err != nil`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *HexTestSuite) TestSuccessCases() {

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
	for _, testCase := range HexSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"Hex\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
