package model

import "fmt"

type Foo struct {
	// Documentation about response Header field.
	Header string `json:"header"`
	Param  string `json:"param"`
	Query  string `json:"query"`
	Foo    string `json:"foo"`
	Incr   int    `json:"incr"`
}

func (foo *Foo) String() string {
	return fmt.Sprintf("%+v", *foo)
}

func NewFoo(header, param, query, foo string, incr int) *Foo {
	return &Foo{
		header,
		param,
		query,
		foo,
		incr,
	}
}
