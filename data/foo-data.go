package data

import (
	"sync/atomic"

	"github.com/nzmprlr/highway/toll"
)

var (
	incr int32
)

type Foo struct {
	toll *toll.Toll
}

func (data *Foo) Incr() (int, error) {
	defer data.toll.Metric(data.toll.Metrics.Data())

	data.toll.Log("foo data Incr")
	data.toll.Log("foo data Incr: %d", incr)

	return int(atomic.AddInt32(&incr, 1)), nil
}

func NewFoo(t *toll.Toll) interface{} {
	return &Foo{
		toll: t,
	}
}
