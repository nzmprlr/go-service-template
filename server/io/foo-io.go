package io

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/nzmprlr/highway/asphalt/errs"
	"github.com/nzmprlr/highway/lane/restserver"

	"{MODULE}/config"
	"{MODULE}/model"
)

type fooRequestBody struct {
	// Documentation about json body field.
	// required:true
	// maxLength:10
	Foo string `json:"foo"`
}

type FooRequest struct {
	// Documentation about header field.
	// in:header
	Header string `json:"header"`
	// in:path
	Param string `json:"param"`
	// in:query
	Query string `json:"query"`

	// Documentation about body.
	// required:true
	// in:body
	Body *fooRequestBody
}

func (foo *FooRequest) Parse(r *http.Request) error {
	errs := errs.New()
	foo.Header = r.Header.Get("header")
	if foo.Header == "" {
		errs.Add(errors.New("header is empty"))
	}

	params := restserver.ParseRequestParams(r)
	foo.Param = params["param"]

	foo.Query = r.URL.Query().Get("query")
	if foo.Query == "" {
		errs.Add(errors.New("query is empty"))
	}

	errs.Add(restserver.ParseRequestBodyJSON(r, foo.Body))

	return errs.Reduce()
}

func (foo *FooRequest) Validate() error {
	errs := errs.New()
	maxLen := config.Get().MaxLenFoo
	if len(foo.Body.Foo) > maxLen {
		errs.Add(fmt.Errorf("foo len error: %d", maxLen))
	}

	return errs.Reduce()
}

func newFooRequest() *FooRequest { return &FooRequest{Body: &fooRequestBody{}} }

type FooResponse struct {
	*model.Foo
	// *model.Foo `json:"foo"`
}

func (foo *FooResponse) Format() error {
	foo.Header = "formatted->" + foo.Header

	return nil
}

func newFooResponse(m *model.Foo) *FooResponse { return &FooResponse{m} }
