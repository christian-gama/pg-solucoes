// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"

	repo "github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

// CourseEnrollment is an autogenerated mock type for the CourseEnrollment type
type CourseEnrollment struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, params
func (_m *CourseEnrollment) Create(ctx context.Context, params repo.CreateCourseEnrollmentParams) (*model.CourseEnrollment, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.CourseEnrollment
	if rf, ok := ret.Get(0).(func(context.Context, repo.CreateCourseEnrollmentParams) *model.CourseEnrollment); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CourseEnrollment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.CreateCourseEnrollmentParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *CourseEnrollment) Delete(ctx context.Context, params repo.DeleteCourseEnrollmentParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.DeleteCourseEnrollmentParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, params, preload
func (_m *CourseEnrollment) FindAll(ctx context.Context, params repo.FindAllCourseEnrollmentParams, preload ...string) (*querying.PaginationOutput[*model.CourseEnrollment], error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *querying.PaginationOutput[*model.CourseEnrollment]
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindAllCourseEnrollmentParams, ...string) *querying.PaginationOutput[*model.CourseEnrollment]); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*model.CourseEnrollment])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindAllCourseEnrollmentParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, params, preload
func (_m *CourseEnrollment) FindOne(ctx context.Context, params repo.FindOneCourseEnrollmentParams, preload ...string) (*model.CourseEnrollment, error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.CourseEnrollment
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindOneCourseEnrollmentParams, ...string) *model.CourseEnrollment); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CourseEnrollment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindOneCourseEnrollmentParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *CourseEnrollment) Update(ctx context.Context, params repo.UpdateCourseEnrollmentParams) (*model.CourseEnrollment, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.CourseEnrollment
	if rf, ok := ret.Get(0).(func(context.Context, repo.UpdateCourseEnrollmentParams) *model.CourseEnrollment); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CourseEnrollment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.UpdateCourseEnrollmentParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCourseEnrollment interface {
	mock.TestingT
	Cleanup(func())
}

// NewCourseEnrollment creates a new instance of CourseEnrollment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCourseEnrollment(t mockConstructorTestingTNewCourseEnrollment) *CourseEnrollment {
	mock := &CourseEnrollment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}