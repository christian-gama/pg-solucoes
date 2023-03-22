// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"

	repo "github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

// Subject is an autogenerated mock type for the Subject type
type Subject struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, params
func (_m *Subject) Create(ctx context.Context, params repo.CreateSubjectParams) (*model.Subject, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, repo.CreateSubjectParams) *model.Subject); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.CreateSubjectParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *Subject) Delete(ctx context.Context, params repo.DeleteSubjectParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.DeleteSubjectParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, params, preload
func (_m *Subject) FindAll(ctx context.Context, params repo.FindAllSubjectParams, preload ...string) (*querying.PaginationOutput[*model.Subject], error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *querying.PaginationOutput[*model.Subject]
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindAllSubjectParams, ...string) *querying.PaginationOutput[*model.Subject]); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*model.Subject])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindAllSubjectParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, params, preload
func (_m *Subject) FindOne(ctx context.Context, params repo.FindOneSubjectParams, preload ...string) (*model.Subject, error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindOneSubjectParams, ...string) *model.Subject); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindOneSubjectParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *Subject) Update(ctx context.Context, params repo.UpdateSubjectParams) (*model.Subject, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, repo.UpdateSubjectParams) *model.Subject); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.UpdateSubjectParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSubject interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubject creates a new instance of Subject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubject(t mockConstructorTestingTNewSubject) *Subject {
	mock := &Subject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}