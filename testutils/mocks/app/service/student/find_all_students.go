// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
)

// FindAllStudents is an autogenerated mock type for the FindAllStudents type
type FindAllStudents struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *FindAllStudents) Handle(ctx context.Context, input *service.FindAllStudentsInput) (*querying.PaginationOutput[*model.Student], error) {
	ret := _m.Called(ctx, input)

	var r0 *querying.PaginationOutput[*model.Student]
	if rf, ok := ret.Get(0).(func(context.Context, *service.FindAllStudentsInput) *querying.PaginationOutput[*model.Student]); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*model.Student])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FindAllStudentsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindAllStudents interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAllStudents creates a new instance of FindAllStudents. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAllStudents(t mockConstructorTestingTNewFindAllStudents) *FindAllStudents {
	mock := &FindAllStudents{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
