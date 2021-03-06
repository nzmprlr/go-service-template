package doc

import "template/server/io"

// swagger:parameters postFooRequest
type getFooRequest struct {
	*io.FooRequest
}

// Documentation about response.
// swagger:response postFooResponse
type getFooResponse struct {
	*io.FooResponse
}
