// Code generated by MockGen. DO NOT EDIT.
// Source: enqueue.go

// Package queue is a generated GoMock package.
package queue

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAsynqEnqueuer is a mock of AsynqEnqueuer interface.
type MockAsynqEnqueuer struct {
	ctrl     *gomock.Controller
	recorder *MockAsynqEnqueuerMockRecorder
}

// MockAsynqEnqueuerMockRecorder is the mock recorder for MockAsynqEnqueuer.
type MockAsynqEnqueuerMockRecorder struct {
	mock *MockAsynqEnqueuer
}

// NewMockAsynqEnqueuer creates a new mock instance.
func NewMockAsynqEnqueuer(ctrl *gomock.Controller) *MockAsynqEnqueuer {
	mock := &MockAsynqEnqueuer{ctrl: ctrl}
	mock.recorder = &MockAsynqEnqueuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsynqEnqueuer) EXPECT() *MockAsynqEnqueuerMockRecorder {
	return m.recorder
}

// Enqueue mocks base method.
func (m *MockAsynqEnqueuer) Enqueue(ctx context.Context, name string, payload any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enqueue", ctx, name, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockAsynqEnqueuerMockRecorder) Enqueue(ctx, name, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockAsynqEnqueuer)(nil).Enqueue), ctx, name, payload)
}
