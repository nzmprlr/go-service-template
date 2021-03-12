package event

import (
	"github.com/nzmprlr/highway"
	"github.com/nzmprlr/highway/toll"

	"{MODULE}/api"
	"{MODULE}/data"
)

type Foo struct {
	toll *toll.Toll

	api.FooData
}

func (foo *Foo) Event() {
	defer highway.Event(foo.toll)
	defer foo.toll.Metric(foo.toll.Metrics.Event())

	foo.toll.Log("foo event Event")

	foo.data().Incr()

	panic("panic test")
}

func (foo *Foo) data() api.FooData {
	return highway.API(foo.toll, foo.FooData, data.NewFoo).(api.FooData)
}

func NewFoo(t *toll.Toll) interface{} {
	return &Foo{
		toll: t.Fork(),
	}
}
