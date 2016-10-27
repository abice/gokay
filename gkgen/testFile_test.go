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
	TestString      string  `valid:"NotEmpty"`
	TestStringPtr   *string `valid:"NotNil"`
	SomeBCP47String string  `valid:"BCP47"`
	SomeInt         *int    `valid:"NoNil"`
}

// SomeOtherStruct
type SomeOtherStruct struct {
	Err     error  `json:"err" valid:"NotNil"`
	Message string `json:"msg"`
}

func (h *SomeImplementation) DoSomething(ctx context.Context, i interface{}) (interface{}, error) {
	return nil, errors.New("Test")
}
