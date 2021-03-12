package doc

import "{MODULE}/server/io"

// swagger:parameters postFooRequest
type postFooRequest struct {
	*io.FooRequest
}

// Documentation about response.
// swagger:response postFooResponse
type postFooResponse struct {
	*io.FooResponse
}
