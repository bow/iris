// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/store.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	internal "github.com/bow/iris/internal"
	gomock "github.com/golang/mock/gomock"
)

// MockFeedStore is a mock of FeedStore interface.
type MockFeedStore struct {
	ctrl     *gomock.Controller
	recorder *MockFeedStoreMockRecorder
}

// MockFeedStoreMockRecorder is the mock recorder for MockFeedStore.
type MockFeedStoreMockRecorder struct {
	mock *MockFeedStore
}

// NewMockFeedStore creates a new mock instance.
func NewMockFeedStore(ctrl *gomock.Controller) *MockFeedStore {
	mock := &MockFeedStore{ctrl: ctrl}
	mock.recorder = &MockFeedStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedStore) EXPECT() *MockFeedStoreMockRecorder {
	return m.recorder
}

// AddFeed mocks base method.
func (m *MockFeedStore) AddFeed(ctx context.Context, feedURL string, title, desc *string, tags []string, isStarred *bool) (*internal.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFeed", ctx, feedURL, title, desc, tags, isStarred)
	ret0, _ := ret[0].(*internal.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFeed indicates an expected call of AddFeed.
func (mr *MockFeedStoreMockRecorder) AddFeed(ctx, feedURL, title, desc, tags, isStarred interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFeed", reflect.TypeOf((*MockFeedStore)(nil).AddFeed), ctx, feedURL, title, desc, tags, isStarred)
}

// DeleteFeeds mocks base method.
func (m *MockFeedStore) DeleteFeeds(ctx context.Context, ids []ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFeeds", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFeeds indicates an expected call of DeleteFeeds.
func (mr *MockFeedStoreMockRecorder) DeleteFeeds(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFeeds", reflect.TypeOf((*MockFeedStore)(nil).DeleteFeeds), ctx, ids)
}

// EditEntries mocks base method.
func (m *MockFeedStore) EditEntries(ctx context.Context, ops []*EntryEditOp) ([]*internal.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditEntries", ctx, ops)
	ret0, _ := ret[0].([]*internal.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditEntries indicates an expected call of EditEntries.
func (mr *MockFeedStoreMockRecorder) EditEntries(ctx, ops interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditEntries", reflect.TypeOf((*MockFeedStore)(nil).EditEntries), ctx, ops)
}

// EditFeeds mocks base method.
func (m *MockFeedStore) EditFeeds(ctx context.Context, ops []*FeedEditOp) ([]*internal.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditFeeds", ctx, ops)
	ret0, _ := ret[0].([]*internal.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditFeeds indicates an expected call of EditFeeds.
func (mr *MockFeedStoreMockRecorder) EditFeeds(ctx, ops interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditFeeds", reflect.TypeOf((*MockFeedStore)(nil).EditFeeds), ctx, ops)
}

// ExportOPML mocks base method.
func (m *MockFeedStore) ExportOPML(ctx context.Context, title *string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportOPML", ctx, title)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportOPML indicates an expected call of ExportOPML.
func (mr *MockFeedStoreMockRecorder) ExportOPML(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportOPML", reflect.TypeOf((*MockFeedStore)(nil).ExportOPML), ctx, title)
}

// GetEntry mocks base method.
func (m *MockFeedStore) GetEntry(ctx context.Context, entryID ID) (*internal.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", ctx, entryID)
	ret0, _ := ret[0].(*internal.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockFeedStoreMockRecorder) GetEntry(ctx, entryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockFeedStore)(nil).GetEntry), ctx, entryID)
}

// GetGlobalStats mocks base method.
func (m *MockFeedStore) GetGlobalStats(ctx context.Context) (*Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalStats", ctx)
	ret0, _ := ret[0].(*Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlobalStats indicates an expected call of GetGlobalStats.
func (mr *MockFeedStoreMockRecorder) GetGlobalStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalStats", reflect.TypeOf((*MockFeedStore)(nil).GetGlobalStats), ctx)
}

// ImportOPML mocks base method.
func (m *MockFeedStore) ImportOPML(ctx context.Context, payload []byte) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportOPML", ctx, payload)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportOPML indicates an expected call of ImportOPML.
func (mr *MockFeedStoreMockRecorder) ImportOPML(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportOPML", reflect.TypeOf((*MockFeedStore)(nil).ImportOPML), ctx, payload)
}

// ListEntries mocks base method.
func (m *MockFeedStore) ListEntries(ctx context.Context, feedID ID) ([]*internal.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", ctx, feedID)
	ret0, _ := ret[0].([]*internal.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockFeedStoreMockRecorder) ListEntries(ctx, feedID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockFeedStore)(nil).ListEntries), ctx, feedID)
}

// ListFeeds mocks base method.
func (m *MockFeedStore) ListFeeds(ctx context.Context) ([]*internal.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeeds", ctx)
	ret0, _ := ret[0].([]*internal.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFeeds indicates an expected call of ListFeeds.
func (mr *MockFeedStoreMockRecorder) ListFeeds(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeeds", reflect.TypeOf((*MockFeedStore)(nil).ListFeeds), ctx)
}

// PullFeeds mocks base method.
func (m *MockFeedStore) PullFeeds(ctx context.Context, feedIDs []ID) <-chan PullResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullFeeds", ctx, feedIDs)
	ret0, _ := ret[0].(<-chan PullResult)
	return ret0
}

// PullFeeds indicates an expected call of PullFeeds.
func (mr *MockFeedStoreMockRecorder) PullFeeds(ctx, feedIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFeeds", reflect.TypeOf((*MockFeedStore)(nil).PullFeeds), ctx, feedIDs)
}

// MockeditableTable is a mock of editableTable interface.
type MockeditableTable struct {
	ctrl     *gomock.Controller
	recorder *MockeditableTableMockRecorder
}

// MockeditableTableMockRecorder is the mock recorder for MockeditableTable.
type MockeditableTableMockRecorder struct {
	mock *MockeditableTable
}

// NewMockeditableTable creates a new mock instance.
func NewMockeditableTable(ctrl *gomock.Controller) *MockeditableTable {
	mock := &MockeditableTable{ctrl: ctrl}
	mock.recorder = &MockeditableTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockeditableTable) EXPECT() *MockeditableTableMockRecorder {
	return m.recorder
}

// errNotFound mocks base method.
func (m *MockeditableTable) errNotFound(id ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "errNotFound", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// errNotFound indicates an expected call of errNotFound.
func (mr *MockeditableTableMockRecorder) errNotFound(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "errNotFound", reflect.TypeOf((*MockeditableTable)(nil).errNotFound), id)
}

// name mocks base method.
func (m *MockeditableTable) name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "name")
	ret0, _ := ret[0].(string)
	return ret0
}

// name indicates an expected call of name.
func (mr *MockeditableTableMockRecorder) name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "name", reflect.TypeOf((*MockeditableTable)(nil).name))
}
