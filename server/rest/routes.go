package rest

import (
	"net/http"

	"github.com/nzmprlr/highway/lane/restserver"
)

func Routes(rest *restserver.Rest) {
	foo := newFoo()
	// swagger:route POST /foo/{param} foo postFooRequest
	// Hello world, foo!
	// Documentation about endpoint.
	// responses:
	//		200: postFooResponse
	rest.AddHandler(http.MethodPost, "/foo/:param", rest.Handler(foo.handleFoo()))
}
