// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
)

// FindOneCourse is an autogenerated mock type for the FindOneCourse type
type FindOneCourse struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *FindOneCourse) Handle(ctx context.Context, input *service.FindOneCourseInput) (*model.Course, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Course
	if rf, ok := ret.Get(0).(func(context.Context, *service.FindOneCourseInput) *model.Course); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FindOneCourseInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindOneCourse interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindOneCourse creates a new instance of FindOneCourse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindOneCourse(t mockConstructorTestingTNewFindOneCourse) *FindOneCourse {
	mock := &FindOneCourse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
