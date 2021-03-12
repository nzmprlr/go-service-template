package io

import (
	"net/http"

	"github.com/nzmprlr/highway/lane/restserver"
	"github.com/nzmprlr/highway/toll"

	"{MODULE}/model"
)

type IO struct {
	toll *toll.Toll
}

func (io *IO) parseRequest(r *http.Request, request restserver.Request) error {
	err := restserver.ParseRequest(r, request)
	if err != nil {
		io.toll.SetStatus(http.StatusBadRequest)
		io.toll.Log(err.Error())
	}

	return err
}

func (io *IO) respondJSON(w http.ResponseWriter, response restserver.Response) error {
	err := restserver.RespondJSON(io.toll, w, response)
	if err != nil {
		io.toll.SetStatus(http.StatusGone)
		io.toll.Log(err.Error())
	}

	return err
}

func (io *IO) ErrorResponse(w http.ResponseWriter, err interface{}) {
	restserver.RespondErrorJSON(io.toll, w, err)
}

func (io *IO) FooRequest(r *http.Request) (*FooRequest, error) {
	defer io.toll.Metric(io.toll.Metrics.IO())

	request := newFooRequest()
	return request, io.parseRequest(r, request)
}

func (io *IO) FooResponse(w http.ResponseWriter, m *model.Foo) error {
	defer io.toll.Metric(io.toll.Metrics.IO(m))

	response := newFooResponse(m)
	return io.respondJSON(w, response)
}

func New(t *toll.Toll) *IO {
	return &IO{
		toll: t,
	}
}
