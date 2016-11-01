package gkgen

import (
	"errors"

	"golang.org/x/net/context"
)

// SomeInterface is awesome.
type SomeInterface interface {
	DoSomething(context.Context, interface{}) (interface{}, error)
}

// SomeImplementation does something
type SomeImplementation struct{}

// NewSomeImplementation creates a Playback Service with default options.
func NewSomeImplementation() SomeInterface {
	return &SomeImplementation{}
}

// SomeStruct is empty because a PlaybackRequest is a GET request.
type SomeStruct struct {
	// URL Deconstructed params
	TestString      string             `valid:"required"`
	TestStringPtr   *string            `valid:"NotNil"`
	SomeBCP47String string             `valid:"BCP47"`
	SomeInt         *int               `valid:"NotNil"`
	RequiredStruct  SomeImplementation `valid:"required"`
	RequiredIFace   SomeInterface      `valid:"required"`
	RequiredChan    chan int           `valid:"required"`
	RequiredBool    bool               `valid:"required"`
	RequiredInt     int                `valid:"required"`
	RequiredInt8    int8               `valid:"required"`
}

// SomeOtherStruct
type SomeOtherStruct struct {
	Err     error  `json:"err" valid:"NotNil"`
	Message string `json:"msg"`
}

func (h *SomeImplementation) DoSomething(ctx context.Context, i interface{}) (interface{}, error) {
	return nil, errors.New("Test")
}
