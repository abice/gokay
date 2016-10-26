package gkgen

import "golang.org/x/net/context"

// DummyInterface
type DummyInterface interface {
	Run(context.Context, interface{}) (interface{}, error)
}
