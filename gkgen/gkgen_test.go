package gkgen

import (
	"fmt"
	"go/scanner"
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
