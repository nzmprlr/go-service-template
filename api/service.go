package api

import "template/model"

type FooService interface {
	Foo(string, string, string, string) (*model.Foo, error)
}

type BarService interface {
	Bar()
}
