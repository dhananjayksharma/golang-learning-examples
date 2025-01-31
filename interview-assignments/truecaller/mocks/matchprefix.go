// Code generated by MockGen. DO NOT EDIT.
// Source: matchprefix.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPrefixFinder is a mock of PrefixFinder interface.
type MockPrefixFinder struct {
	ctrl     *gomock.Controller
	recorder *MockPrefixFinderMockRecorder
}

// MockPrefixFinderMockRecorder is the mock recorder for MockPrefixFinder.
type MockPrefixFinderMockRecorder struct {
	mock *MockPrefixFinder
}

// NewMockPrefixFinder creates a new mock instance.
func NewMockPrefixFinder(ctrl *gomock.Controller) *MockPrefixFinder {
	mock := &MockPrefixFinder{ctrl: ctrl}
	mock.recorder = &MockPrefixFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrefixFinder) EXPECT() *MockPrefixFinderMockRecorder {
	return m.recorder
}

// FindPrefix mocks base method.
func (m *MockPrefixFinder) FindPrefix(input string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrefix", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindPrefix indicates an expected call of FindPrefix.
func (mr *MockPrefixFinderMockRecorder) FindPrefix(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrefix", reflect.TypeOf((*MockPrefixFinder)(nil).FindPrefix), input)
}
