// Code generated by MockGen. DO NOT EDIT.
// Source: internal/internal.go

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	store "github.com/bow/courier/internal/store"
	gomock "github.com/golang/mock/gomock"
	gofeed "github.com/mmcdole/gofeed"
)

// MockFeedParser is a mock of FeedParser interface.
type MockFeedParser struct {
	ctrl     *gomock.Controller
	recorder *MockFeedParserMockRecorder
}

// MockFeedParserMockRecorder is the mock recorder for MockFeedParser.
type MockFeedParserMockRecorder struct {
	mock *MockFeedParser
}

// NewMockFeedParser creates a new mock instance.
func NewMockFeedParser(ctrl *gomock.Controller) *MockFeedParser {
	mock := &MockFeedParser{ctrl: ctrl}
	mock.recorder = &MockFeedParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedParser) EXPECT() *MockFeedParserMockRecorder {
	return m.recorder
}

// ParseURLWithContext mocks base method.
func (m *MockFeedParser) ParseURLWithContext(feedURL string, ctx context.Context) (*gofeed.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseURLWithContext", feedURL, ctx)
	ret0, _ := ret[0].(*gofeed.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseURLWithContext indicates an expected call of ParseURLWithContext.
func (mr *MockFeedParserMockRecorder) ParseURLWithContext(feedURL, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseURLWithContext", reflect.TypeOf((*MockFeedParser)(nil).ParseURLWithContext), feedURL, ctx)
}

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
func (m *MockFeedStore) AddFeed(ctx context.Context, feed *gofeed.Feed, title, desc *string, categories []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFeed", ctx, feed, title, desc, categories)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFeed indicates an expected call of AddFeed.
func (mr *MockFeedStoreMockRecorder) AddFeed(ctx, feed, title, desc, categories interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFeed", reflect.TypeOf((*MockFeedStore)(nil).AddFeed), ctx, feed, title, desc, categories)
}

// ListFeeds mocks base method.
func (m *MockFeedStore) ListFeeds(ctx context.Context) ([]*store.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeeds", ctx)
	ret0, _ := ret[0].([]*store.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFeeds indicates an expected call of ListFeeds.
func (mr *MockFeedStoreMockRecorder) ListFeeds(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeeds", reflect.TypeOf((*MockFeedStore)(nil).ListFeeds), ctx)
}
