package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nzmprlr/highway/toll"
	"github.com/stretchr/testify/assert"

	mock_api "template/mock"
)

func TestFoo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	toll := toll.New()
	service := NewFoo(toll).(*Foo)
	fooData := mock_api.NewMockFooData(mockCtrl)
	service.FooData = fooData

	fooData.EXPECT().Incr().Return(123, nil).Times(1)
	fooData.EXPECT().Incr().Return(0, errors.New("err")).Times(1)

	model, err := service.Foo("h", "p", "q", "f")
	assert.NoError(t, err)
	assert.Equal(t, model.Incr, 123)

	model, err = service.Foo("h", "p", "q", "f")
	assert.EqualError(t, err, "err")
}
