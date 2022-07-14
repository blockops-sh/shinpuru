// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// TimeProvider is an autogenerated mock type for the Provider type
type TimeProvider struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *TimeProvider) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

type mockConstructorTestingTNewTimeProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimeProvider creates a new instance of TimeProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimeProvider(t mockConstructorTestingTNewTimeProvider) *TimeProvider {
	mock := &TimeProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
