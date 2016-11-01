package gkgen

import (
	"fmt"
	"go/scanner"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testInput     = `testFile_test.go`
	testNoStructs = `nostructs_test.go`
	testExample   = `example.test`
)

// GkgenTestSuite
type GkgenTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *GkgenTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestGkgenTestSuite(t *testing.T) {
	suite.Run(t, new(GkgenTestSuite))
}

// TestInputFile
func (s *GkgenTestSuite) TestInputFile() {
	g := NewGenerator()
	// Parse the file given in arguments
	imported, err := g.GenerateFromFile(testInput)
	if eList, ok := err.(scanner.ErrorList); ok {
		s.FailNow("Error generating code: " + eList.Error())
	}
	s.Nil(err, "Error generating formatted code")
	fmt.Println(string(imported))
}

// TestNoStructInputFile
func (s *GkgenTestSuite) TestNoStructFile() {
	g := NewGenerator()
	// Parse the file given in arguments
	imported, err := g.GenerateFromFile(testNoStructs)
	s.Nil(err, "Error generating formatted code")
	fmt.Println(string(imported))
}

// TestNoFile
func (s *GkgenTestSuite) TestNoFile() {
	g := NewGenerator()
	// Parse the file given in arguments
	_, err := g.GenerateFromFile("")
	s.NotNil(err, "Error generating formatted code")
}

// TestExampleFile
func (s *GkgenTestSuite) TestExampleFile() {
	g := NewGenerator()
	// Parse the file given in arguments
	imported, err := g.GenerateFromFile(testExample)
	s.Nil(err, "Error generating formatted code")
	fmt.Println(string(imported))
}

var result bool

// BenchmarkReflection is a quick test to see how much of an impact reflection has
// in the performance of an application
func BenchmarkReflection(b *testing.B) {
	match := false
	for x := 0; x < b.N; x++ {
		zeroInt := reflect.Zero(reflect.ValueOf(x).Type())
		if reflect.ValueOf(x).Interface() == zeroInt.Interface() {
			match = true
		}
		match = false
	}
	result = match
}

// BenchmarkEmptyVar is a quick benchmark to determine performance of just using an empty var
// for zero value comparison
func BenchmarkEmptyVar(b *testing.B) {
	match := false
	for x := 0; x < b.N; x++ {
		var zeroInt int
		if x == zeroInt {
			match = true
		}
		match = false
	}
	result = match
}
