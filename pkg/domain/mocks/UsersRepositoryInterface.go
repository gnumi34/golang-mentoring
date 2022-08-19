// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	uservice "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
	mock "github.com/stretchr/testify/mock"
)

// UsersRepositoryInterface is an autogenerated mock type for the UsersRepositoryInterface type
type UsersRepositoryInterface struct {
	mock.Mock
}

// AddUsers provides a mock function with given fields: ctx, userDomain
func (_m *UsersRepositoryInterface) AddUser(ctx context.Context, userDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	ret := _m.Called(ctx, userDomain)

	var r0 uservice.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, uservice.UsersDomain) uservice.UsersDomain); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(uservice.UsersDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uservice.UsersDomain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUsers provides a mock function with given fields: CTX, id
func (_m *UsersRepositoryInterface) DeleteUser(CTX context.Context, id string) error {
	ret := _m.Called(CTX, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(CTX, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, userDomain
func (_m *UsersRepositoryInterface) GetUser(ctx context.Context, userDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	ret := _m.Called(ctx, userDomain)

	var r0 uservice.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, uservice.UsersDomain) uservice.UsersDomain); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(uservice.UsersDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uservice.UsersDomain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUsers provides a mock function with given fields: ctx, userUpdateDomain
func (_m *UsersRepositoryInterface) UpdateUser(ctx context.Context, userUpdateDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	ret := _m.Called(ctx, userUpdateDomain)

	var r0 uservice.UsersDomain
	if rf, ok := ret.Get(0).(func(context.Context, uservice.UsersDomain) uservice.UsersDomain); ok {
		r0 = rf(ctx, userUpdateDomain)
	} else {
		r0 = ret.Get(0).(uservice.UsersDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uservice.UsersDomain) error); ok {
		r1 = rf(ctx, userUpdateDomain)
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