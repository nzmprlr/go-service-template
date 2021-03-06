package service

import (
	"github.com/nzmprlr/highway/toll"
)

type Bar struct {
	toll *toll.Toll
}

func (service *Bar) Bar() {
	defer service.toll.Metric(service.toll.Metrics.Service())

	service.toll.Label("cache", true)
	service.toll.Log("bar service Bar")
}

func NewBar(t *toll.Toll) interface{} {
	return &Bar{
		toll: t,
	}
}
