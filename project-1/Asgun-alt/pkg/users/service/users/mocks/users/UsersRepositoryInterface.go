// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	users "golang-mentoring/project-1/Asgun-alt/pkg/domain/users"

	mock "github.com/stretchr/testify/mock"
)

// UsersRepositoryInterface is an autogenerated mock type for the UsersRepositoryInterface type
type UsersRepositoryInterface struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: ctx, req
func (_m *UsersRepositoryInterface) AddUser(ctx context.Context, req *users.UsersDomain) (*users.UsersDomain, error) {
	ret := _m.Called(ctx, req)

	var r0 *users.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, *users.UsersDomain) *users.UsersDomain); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UsersDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.UsersDomain) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UsersRepositoryInterface) DeleteUser(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserByID provides a mock function with given fields: ctx, id
func (_m *UsersRepositoryInterface) FindUserByID(ctx context.Context, id uint) (*users.UsersDomain, error) {
	ret := _m.Called(ctx, id)

	var r0 *users.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, uint) *users.UsersDomain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UsersDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, req
func (_m *UsersRepositoryInterface) GetUser(ctx context.Context, req *users.UsersDomain) (*users.UsersDomain, error) {
	ret := _m.Called(ctx, req)

	var r0 *users.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, *users.UsersDomain) *users.UsersDomain); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UsersDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.UsersDomain) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, req
func (_m *UsersRepositoryInterface) UpdateUser(ctx context.Context, req *users.UsersDomain) (*users.UsersDomain, error) {
	ret := _m.Called(ctx, req)

	var r0 *users.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, *users.UsersDomain) *users.UsersDomain); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UsersDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.UsersDomain) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsersRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersRepositoryInterface creates a new instance of UsersRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepositoryInterface(t mockConstructorTestingTNewUsersRepositoryInterface) *UsersRepositoryInterface {
	mock := &UsersRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
