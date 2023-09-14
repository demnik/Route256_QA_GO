// Code generated by MockGen. DO NOT EDIT.
// Source: ./event.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ozonmp/act-device-api/internal/model"
)

// MockEventRepo is a mock of EventRepo interface.
type MockEventRepo struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepoMockRecorder
}

// MockEventRepoMockRecorder is the mock recorder for MockEventRepo.
type MockEventRepoMockRecorder struct {
	mock *MockEventRepo
}

// NewMockEventRepo creates a new mock instance.
func NewMockEventRepo(ctrl *gomock.Controller) *MockEventRepo {
	mock := &MockEventRepo{ctrl: ctrl}
	mock.recorder = &MockEventRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepo) EXPECT() *MockEventRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockEventRepo) Add(ctx context.Context, event *model.DeviceEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockEventRepoMockRecorder) Add(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEventRepo)(nil).Add), ctx, event)
}

// Lock mocks base method.
func (m *MockEventRepo) Lock(ctx context.Context, n uint64) ([]model.DeviceEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, n)
	ret0, _ := ret[0].([]model.DeviceEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockEventRepoMockRecorder) Lock(ctx, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockEventRepo)(nil).Lock), ctx, n)
}

// Remove mocks base method.
func (m *MockEventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, eventIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockEventRepoMockRecorder) Remove(ctx, eventIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockEventRepo)(nil).Remove), ctx, eventIDs)
}

// Unlock mocks base method.
func (m *MockEventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", ctx, eventIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock.
func (mr *MockEventRepoMockRecorder) Unlock(ctx, eventIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockEventRepo)(nil).Unlock), ctx, eventIDs)
}