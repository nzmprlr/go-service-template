package service

import (
	"github.com/nzmprlr/highway"
	"github.com/nzmprlr/highway/toll"

	"{MODULE}/api"
	"{MODULE}/call"
	"{MODULE}/data"
	"{MODULE}/event"
	"{MODULE}/model"
)

type Foo struct {
	toll *toll.Toll

	api.FooEvent
	api.FooData
	api.FooCall
	api.BarService
}

func (service *Foo) Foo(header, param, query, foo string) (*model.Foo, error) {
	defer service.toll.Metric(service.toll.Metrics.Service(header, param, query, foo))

	go service.event().Event()

	service.toll.Log("foo service %s", "Foo")

	service.bar().Bar()

	incr, err := service.data().Incr()
	if err != nil {
		return nil, err
	}

	service.call().Call()

	return model.NewFoo(header, param, query, foo, incr), nil
}

func (service *Foo) data() api.FooData {
	return highway.API(service.toll, service.FooData, data.NewFoo).(api.FooData)
}

func (service *Foo) call() api.FooCall {
	return highway.API(service.toll, service.FooCall, call.NewFoo).(api.FooCall)
}

func (service *Foo) event() api.FooEvent {
	return highway.API(service.toll, service.FooEvent, event.NewFoo).(api.FooEvent)
}

func (service *Foo) bar() api.BarService {
	return highway.API(service.toll, service.BarService, NewBar).(api.BarService)
}

func NewFoo(t *toll.Toll) interface{} {
	return &Foo{
		toll: t,
	}
}
