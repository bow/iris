// Code generated by MockGen. DO NOT EDIT.
// Source: internal/datastore/datastore.go
//
// Generated by this command:
//
//	mockgen -source=internal/datastore/datastore.go -package=server Datastore
//

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"

	datastore "github.com/bow/neon/internal/datastore"
	entity "github.com/bow/neon/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockDatastore is a mock of Datastore interface.
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore.
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance.
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// AddFeed mocks base method.
func (m *MockDatastore) AddFeed(ctx context.Context, feedURL string, title, desc *string, tags []string, isStarred *bool) (*entity.Feed, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFeed", ctx, feedURL, title, desc, tags, isStarred)
	ret0, _ := ret[0].(*entity.Feed)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddFeed indicates an expected call of AddFeed.
func (mr *MockDatastoreMockRecorder) AddFeed(ctx, feedURL, title, desc, tags, isStarred any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFeed", reflect.TypeOf((*MockDatastore)(nil).AddFeed), ctx, feedURL, title, desc, tags, isStarred)
}

// DeleteFeeds mocks base method.
func (m *MockDatastore) DeleteFeeds(ctx context.Context, ids []entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFeeds", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFeeds indicates an expected call of DeleteFeeds.
func (mr *MockDatastoreMockRecorder) DeleteFeeds(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFeeds", reflect.TypeOf((*MockDatastore)(nil).DeleteFeeds), ctx, ids)
}

// EditEntries mocks base method.
func (m *MockDatastore) EditEntries(ctx context.Context, ops []*entity.EntryEditOp) ([]*entity.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditEntries", ctx, ops)
	ret0, _ := ret[0].([]*entity.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditEntries indicates an expected call of EditEntries.
func (mr *MockDatastoreMockRecorder) EditEntries(ctx, ops any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditEntries", reflect.TypeOf((*MockDatastore)(nil).EditEntries), ctx, ops)
}

// EditFeeds mocks base method.
func (m *MockDatastore) EditFeeds(ctx context.Context, ops []*entity.FeedEditOp) ([]*entity.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditFeeds", ctx, ops)
	ret0, _ := ret[0].([]*entity.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditFeeds indicates an expected call of EditFeeds.
func (mr *MockDatastoreMockRecorder) EditFeeds(ctx, ops any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditFeeds", reflect.TypeOf((*MockDatastore)(nil).EditFeeds), ctx, ops)
}

// ExportSubscription mocks base method.
func (m *MockDatastore) ExportSubscription(ctx context.Context, title *string) (*entity.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSubscription", ctx, title)
	ret0, _ := ret[0].(*entity.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportSubscription indicates an expected call of ExportSubscription.
func (mr *MockDatastoreMockRecorder) ExportSubscription(ctx, title any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSubscription", reflect.TypeOf((*MockDatastore)(nil).ExportSubscription), ctx, title)
}

// GetEntry mocks base method.
func (m *MockDatastore) GetEntry(ctx context.Context, id entity.ID) (*entity.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", ctx, id)
	ret0, _ := ret[0].(*entity.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockDatastoreMockRecorder) GetEntry(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockDatastore)(nil).GetEntry), ctx, id)
}

// GetGlobalStats mocks base method.
func (m *MockDatastore) GetGlobalStats(ctx context.Context) (*entity.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalStats", ctx)
	ret0, _ := ret[0].(*entity.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlobalStats indicates an expected call of GetGlobalStats.
func (mr *MockDatastoreMockRecorder) GetGlobalStats(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalStats", reflect.TypeOf((*MockDatastore)(nil).GetGlobalStats), ctx)
}

// ImportSubscription mocks base method.
func (m *MockDatastore) ImportSubscription(ctx context.Context, sub *entity.Subscription) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportSubscription", ctx, sub)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportSubscription indicates an expected call of ImportSubscription.
func (mr *MockDatastoreMockRecorder) ImportSubscription(ctx, sub any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportSubscription", reflect.TypeOf((*MockDatastore)(nil).ImportSubscription), ctx, sub)
}

// ListEntries mocks base method.
func (m *MockDatastore) ListEntries(ctx context.Context, feedIDs []entity.ID, isBookmarked *bool) ([]*entity.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", ctx, feedIDs, isBookmarked)
	ret0, _ := ret[0].([]*entity.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockDatastoreMockRecorder) ListEntries(ctx, feedIDs, isBookmarked any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockDatastore)(nil).ListEntries), ctx, feedIDs, isBookmarked)
}

// ListFeeds mocks base method.
func (m *MockDatastore) ListFeeds(ctx context.Context, maxEntriesPerFeed *uint32) ([]*entity.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeeds", ctx, maxEntriesPerFeed)
	ret0, _ := ret[0].([]*entity.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFeeds indicates an expected call of ListFeeds.
func (mr *MockDatastoreMockRecorder) ListFeeds(ctx, maxEntriesPerFeed any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeeds", reflect.TypeOf((*MockDatastore)(nil).ListFeeds), ctx, maxEntriesPerFeed)
}

// PullFeeds mocks base method.
func (m *MockDatastore) PullFeeds(ctx context.Context, ids []entity.ID, entryReadStatus *bool, maxEntriesPerFeed *uint32) <-chan entity.PullResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullFeeds", ctx, ids, entryReadStatus, maxEntriesPerFeed)
	ret0, _ := ret[0].(<-chan entity.PullResult)
	return ret0
}

// PullFeeds indicates an expected call of PullFeeds.
func (mr *MockDatastoreMockRecorder) PullFeeds(ctx, ids, entryReadStatus, maxEntriesPerFeed any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFeeds", reflect.TypeOf((*MockDatastore)(nil).PullFeeds), ctx, ids, entryReadStatus, maxEntriesPerFeed)
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
func (m *MockeditableTable) errNotFound(id datastore.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "errNotFound", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// errNotFound indicates an expected call of errNotFound.
func (mr *MockeditableTableMockRecorder) errNotFound(id any) *gomock.Call {
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
