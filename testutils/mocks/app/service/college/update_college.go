// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
)

// UpdateCollege is an autogenerated mock type for the UpdateCollege type
type UpdateCollege struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *UpdateCollege) Handle(ctx context.Context, input *service.UpdateCollegeInput) (*model.College, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.College
	if rf, ok := ret.Get(0).(func(context.Context, *service.UpdateCollegeInput) *model.College); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.College)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.UpdateCollegeInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUpdateCollege interface {
	mock.TestingT
	Cleanup(func())
}

// NewUpdateCollege creates a new instance of UpdateCollege. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUpdateCollege(t mockConstructorTestingTNewUpdateCollege) *UpdateCollege {
	mock := &UpdateCollege{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
