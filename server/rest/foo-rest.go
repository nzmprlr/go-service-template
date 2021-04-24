package rest

import (
	"net/http"

	"github.com/nzmprlr/highway"
	"github.com/nzmprlr/highway/lane/restserver"
	"github.com/nzmprlr/highway/toll"

	"{MODULE}/api"
	"{MODULE}/server/io"
	"{MODULE}/service"
)

type foo struct {
	api.FooService
}

func (handler *foo) handleFoo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toll := restserver.ParseToll(r)
		defer toll.Metric(toll.Metrics.Controller())

		io := io.New(toll)
		request, err := io.FooRequest(r)
		if err != nil {
			io.ErrorResponse(w, err)
			return
		}

		service := handler.service(toll)
		foo, err := service.Foo(request.Header, request.Param, request.Query, request.Body.Foo)
		if err != nil {
			io.ErrorResponse(w, err)
			return
		}

		err = io.FooResponse(w, foo)
		if err != nil {
			io.ErrorResponse(w, err)
		}
	}
}


func (handler *foo) service(toll *toll.Toll) api.FooService {
	return highway.API(toll, handler.FooService, service.NewFoo).(api.FooService)
}

func newFoo() *foo {
	return &foo{}
}
