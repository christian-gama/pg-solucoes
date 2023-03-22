// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	mock "github.com/stretchr/testify/mock"
)

// FindOneCollege is an autogenerated mock type for the FindOneCollege type
type FindOneCollege struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *FindOneCollege) Handle(ctx context.Context, input *service.FindOneCollegeInput) (*service.FindOneCollegeOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *service.FindOneCollegeOutput
	if rf, ok := ret.Get(0).(func(context.Context, *service.FindOneCollegeInput) *service.FindOneCollegeOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.FindOneCollegeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FindOneCollegeInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindOneCollege interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindOneCollege creates a new instance of FindOneCollege. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindOneCollege(t mockConstructorTestingTNewFindOneCollege) *FindOneCollege {
	mock := &FindOneCollege{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}