package rest

import (
	"net/http"

	"github.com/nzmprlr/highway/lane/restserver"

	"template/api"
	"template/server/io"
	"template/service"
)

func handleFoo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toll := restserver.ParseToll(r)
		defer toll.Metric(toll.Metrics.Controller())

		io := io.New(toll)
		request, err := io.FooRequest(r)
		if err != nil {
			io.ErrorResponse(w, err)
			return
		}

		service := service.NewFoo(toll).(api.FooService)
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
