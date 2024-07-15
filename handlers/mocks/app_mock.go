// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	dto "github.com/LiFeAiR/ozon_sea_fight/dto"

	mock "github.com/stretchr/testify/mock"
)

// AppMock is an autogenerated mock type for the AppMock type
type AppMock struct {
	mock.Mock
}

// Clear provides a mock function with given fields:
func (_m *AppMock) Clear() {
	_m.Called()
}

// CreateMatrix provides a mock function with given fields: maxIndex
func (_m *AppMock) CreateMatrix(maxIndex int) {
	_m.Called(maxIndex)
}

// CreateShips provides a mock function with given fields: coordinates
func (_m *AppMock) CreateShips(coordinates string) error {
	ret := _m.Called(coordinates)

	if len(ret) == 0 {
		panic("no return value specified for CreateShips")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(coordinates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFightMatrix provides a mock function with given fields:
func (_m *AppMock) GetFightMatrix() map[string]map[string]bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFightMatrix")
	}

	var r0 map[string]map[string]bool
	if rf, ok := ret.Get(0).(func() map[string]map[string]bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]bool)
		}
	}

	return r0
}

// MakeShot provides a mock function with given fields: coordinates
func (_m *AppMock) MakeShot(coordinates string) (*dto.Shot, error) {
	ret := _m.Called(coordinates)

	if len(ret) == 0 {
		panic("no return value specified for MakeShot")
	}

	var r0 *dto.Shot
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*dto.Shot, error)); ok {
		return rf(coordinates)
	}
	if rf, ok := ret.Get(0).(func(string) *dto.Shot); ok {
		r0 = rf(coordinates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Shot)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(coordinates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShipsCreated provides a mock function with given fields:
func (_m *AppMock) ShipsCreated() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ShipsCreated")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// State provides a mock function with given fields:
func (_m *AppMock) State() dto.State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for State")
	}

	var r0 dto.State
	if rf, ok := ret.Get(0).(func() dto.State); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dto.State)
	}

	return r0
}

// NewAppMock creates a new instance of AppMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppMock {
	mock := &AppMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}