package api

import "{MODULE}/model"

type FooService interface {
	Foo(string, string, string, string) (*model.Foo, error)
}

type BarService interface {
	Bar()
}
