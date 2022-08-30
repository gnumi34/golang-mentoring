// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/albertafriadii/tree/fix/albertafriadii/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecaseInterface is an autogenerated mock type for the UserUsecaseInterface type
type UserUsecaseInterface struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *UserUsecaseInterface) CreateUser(ctx context.Context, u domain.Users) (domain.Users, error) {
	ret := _m.Called(ctx, u)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(context.Context, domain.Users) domain.Users); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Users) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, user_id
func (_m *UserUsecaseInterface) DeleteUser(ctx context.Context, user_id string) error {
	ret := _m.Called(ctx, user_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, user_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, u
func (_m *UserUsecaseInterface) GetUser(ctx context.Context, u domain.Users) (domain.Users, error) {
	ret := _m.Called(ctx, u)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(context.Context, domain.Users) domain.Users); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Users) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: ctx, u
func (_m *UserUsecaseInterface) LoginUser(ctx context.Context, u domain.Users) (bool, error) {
	ret := _m.Called(ctx, u)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, domain.Users) bool); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Users) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, u, user_id
func (_m *UserUsecaseInterface) UpdateUser(ctx context.Context, u domain.Users, user_id string) (domain.Users, error) {
	ret := _m.Called(ctx, u, user_id)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(context.Context, domain.Users, string) domain.Users); ok {
		r0 = rf(ctx, u, user_id)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Users, string) error); ok {
		r1 = rf(ctx, u, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecaseInterface creates a new instance of UserUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecaseInterface(t mockConstructorTestingTNewUserUsecaseInterface) *UserUsecaseInterface {
	mock := &UserUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
