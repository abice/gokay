package gkgen

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// UUIDTestSuite
type UUIDTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *UUIDTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestUUIDTestSuite(t *testing.T) {
	suite.Run(t, new(UUIDTestSuite))
}

var UUIDErrorCases = []struct {
	inType      string
	errorString string
}{
	{"int", `UUID is not valid on field 'TestField int'`},
	{"int8", `UUID is not valid on field 'TestField int8'`},
	{"int16", `UUID is not valid on field 'TestField int16'`},
	{"int32", `UUID is not valid on field 'TestField int32'`},
	{"int64", `UUID is not valid on field 'TestField int64'`},
	{"uint", `UUID is not valid on field 'TestField uint'`},
	{"uint8", `UUID is not valid on field 'TestField uint8'`},
	{"uint16", `UUID is not valid on field 'TestField uint16'`},
	{"uint32", `UUID is not valid on field 'TestField uint32'`},
	{"uint64", `UUID is not valid on field 'TestField uint64'`},
	{"bool", `UUID is not valid on field 'TestField bool'`},
	{"float", `UUID is not valid on field 'TestField float'`},
	{"float32", `UUID is not valid on field 'TestField float32'`},
	{"float64", `UUID is not valid on field 'TestField float64'`},
	{"complex64", `UUID is not valid on field 'TestField complex64'`},
	{"complex128", `UUID is not valid on field 'TestField complex128'`},
	{"SomeOtherStruct", `UUID is not valid on field 'TestField SomeOtherStruct'`},
	{"*SomeOtherStruct", `UUID is not valid on field 'TestField *SomeOtherStruct'`},
	{"SomeInterface", `UUID is not valid on field 'TestField SomeInterface'`},
}

func (s *UUIDTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range UUIDErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"UUID\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error for '%s' missing", testCase.inType)
		if err != nil {
			s.Contains(err.Error(), testCase.errorString)
		}
	}
}

var UUIDSuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"*string", `gokay.IsUUID(s.TestField); err != nil`},
	{"string", `gokay.IsUUID(&s.TestField); err != nil`},
	{"[]string", `gokay.IsUUID(&singleTestField); err != nil`},
	{"[]*string", `gokay.IsUUID(singleTestField); err != nil`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *UUIDTestSuite) TestSuccessCases() {

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
	for _, testCase := range UUIDSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"UUID\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
