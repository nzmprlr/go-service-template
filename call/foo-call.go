package call

import (
	"github.com/nzmprlr/highway/toll"
)

type Foo struct {
	toll *toll.Toll
}

func (call *Foo) Call() {
	defer call.toll.Metric(call.toll.Metrics.Call())

	call.toll.Log("foo call Call")
}

func NewFoo(t *toll.Toll) interface{} {
	return &Foo{
		toll: t,
	}
}
