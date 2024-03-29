// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// IsActive mocks base method.
func (m *MockDriver) IsActive(arg0 string, arg1 map[string]interface{}) *bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	return ret0
}

// IsActive indicates an expected call of IsActive.
func (mr *MockDriverMockRecorder) IsActive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockDriver)(nil).IsActive), arg0, arg1)
}
