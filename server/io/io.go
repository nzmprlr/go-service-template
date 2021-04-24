package io

import (
	"net/http"

	"github.com/nzmprlr/highway/lane/restserver"
	"github.com/nzmprlr/highway/toll"
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

func New(t *toll.Toll) *IO {
	return &IO{
		toll: t,
	}
}
