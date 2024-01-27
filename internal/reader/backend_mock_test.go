// Code generated by MockGen. DO NOT EDIT.
// Source: internal/reader/backend/backend.go
//
// Generated by this command:
//
//	mockgen -source=internal/reader/backend/backend.go -package=reader Backend
//

// Package reader is a generated GoMock package.
package reader

import (
	context "context"
	reflect "reflect"

	entity "github.com/bow/neon/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// GetAllFeedsF mocks base method.
func (m *MockBackend) GetAllFeedsF(arg0 context.Context) func() ([]*entity.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFeedsF", arg0)
	ret0, _ := ret[0].(func() ([]*entity.Feed, error))
	return ret0
}

// GetAllFeedsF indicates an expected call of GetAllFeedsF.
func (mr *MockBackendMockRecorder) GetAllFeedsF(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFeedsF", reflect.TypeOf((*MockBackend)(nil).GetAllFeedsF), arg0)
}

// GetStatsF mocks base method.
func (m *MockBackend) GetStatsF(arg0 context.Context) func() (*entity.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatsF", arg0)
	ret0, _ := ret[0].(func() (*entity.Stats, error))
	return ret0
}

// GetStatsF indicates an expected call of GetStatsF.
func (mr *MockBackendMockRecorder) GetStatsF(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatsF", reflect.TypeOf((*MockBackend)(nil).GetStatsF), arg0)
}

// PullFeedsF mocks base method.
func (m *MockBackend) PullFeedsF(arg0 context.Context, arg1 []entity.ID) func() (<-chan entity.PullResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullFeedsF", arg0, arg1)
	ret0, _ := ret[0].(func() (<-chan entity.PullResult, error))
	return ret0
}

// PullFeedsF indicates an expected call of PullFeedsF.
func (mr *MockBackendMockRecorder) PullFeedsF(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFeedsF", reflect.TypeOf((*MockBackend)(nil).PullFeedsF), arg0, arg1)
}

// String mocks base method.
func (m *MockBackend) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockBackendMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBackend)(nil).String))
}
