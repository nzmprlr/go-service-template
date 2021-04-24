package rest

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"{MODULE}/config"
	mock_api "{MODULE}/mock"
)

func TestFoo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	handler := newFoo()
	fooService := mock_api.NewMockFooService(mockCtrl)
	handler.FooService = fooService

	fooService.EXPECT().Foo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("err")).Times(1)

	r := httptest.NewRequest(http.MethodGet, "/foo/param", bytes.NewBuffer([]byte(`{"foo":"foo"}`)))
	r.Header.Set("header", "h")
	q := r.URL.Query()
	q.Add("query", "q")
	r.URL.RawQuery = q.Encode()

	config.Get().MaxLenFoo = 5

	w := httptest.NewRecorder()
	handler.handleFoo()(w, r)

	assert.Equal(t, `{"error":"err"}`, strings.TrimSpace(w.Body.String()))
}
