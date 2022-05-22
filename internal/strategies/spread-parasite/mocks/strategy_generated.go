// Code generated by MockGen. DO NOT EDIT.
// Source: strategy.go

// Package spreadparasitemocks is a generated GoMock package.
package spreadparasitemocks

import (
	context "context"
	reflect "reflect"

	tinkoffinvest "github.com/Antonboom/tinkoff-invest-robot-contest-2022/internal/clients/tinkoffinvest"
	toolscache "github.com/Antonboom/tinkoff-invest-robot-contest-2022/internal/services/tools-cache"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockOrderPlacer is a mock of OrderPlacer interface.
type MockOrderPlacer struct {
	ctrl     *gomock.Controller
	recorder *MockOrderPlacerMockRecorder
}

// MockOrderPlacerMockRecorder is the mock recorder for MockOrderPlacer.
type MockOrderPlacerMockRecorder struct {
	mock *MockOrderPlacer
}

// NewMockOrderPlacer creates a new mock instance.
func NewMockOrderPlacer(ctrl *gomock.Controller) *MockOrderPlacer {
	mock := &MockOrderPlacer{ctrl: ctrl}
	mock.recorder = &MockOrderPlacerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderPlacer) EXPECT() *MockOrderPlacerMockRecorder {
	return m.recorder
}

// CancelOrder mocks base method.
func (m *MockOrderPlacer) CancelOrder(ctx context.Context, orderID tinkoffinvest.OrderID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockOrderPlacerMockRecorder) CancelOrder(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockOrderPlacer)(nil).CancelOrder), ctx, orderID)
}

// GetOrderBook mocks base method.
func (m *MockOrderPlacer) GetOrderBook(ctx context.Context, req tinkoffinvest.OrderBookRequest) (*tinkoffinvest.OrderBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderBook", ctx, req)
	ret0, _ := ret[0].(*tinkoffinvest.OrderBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderBook indicates an expected call of GetOrderBook.
func (mr *MockOrderPlacerMockRecorder) GetOrderBook(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderBook", reflect.TypeOf((*MockOrderPlacer)(nil).GetOrderBook), ctx, req)
}

// GetOrderState mocks base method.
func (m *MockOrderPlacer) GetOrderState(ctx context.Context, arg1 tinkoffinvest.AccountID, arg2 tinkoffinvest.OrderID) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderState", ctx, arg1, arg2)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderState indicates an expected call of GetOrderState.
func (mr *MockOrderPlacerMockRecorder) GetOrderState(ctx, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderState", reflect.TypeOf((*MockOrderPlacer)(nil).GetOrderState), ctx, arg1, arg2)
}

// GetTradeAvailableShares mocks base method.
func (m *MockOrderPlacer) GetTradeAvailableShares(ctx context.Context) ([]tinkoffinvest.Instrument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTradeAvailableShares", ctx)
	ret0, _ := ret[0].([]tinkoffinvest.Instrument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTradeAvailableShares indicates an expected call of GetTradeAvailableShares.
func (mr *MockOrderPlacerMockRecorder) GetTradeAvailableShares(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTradeAvailableShares", reflect.TypeOf((*MockOrderPlacer)(nil).GetTradeAvailableShares), ctx)
}

// PlaceLimitBuyOrder mocks base method.
func (m *MockOrderPlacer) PlaceLimitBuyOrder(ctx context.Context, request tinkoffinvest.PlaceOrderRequest) (tinkoffinvest.OrderID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceLimitBuyOrder", ctx, request)
	ret0, _ := ret[0].(tinkoffinvest.OrderID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlaceLimitBuyOrder indicates an expected call of PlaceLimitBuyOrder.
func (mr *MockOrderPlacerMockRecorder) PlaceLimitBuyOrder(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceLimitBuyOrder", reflect.TypeOf((*MockOrderPlacer)(nil).PlaceLimitBuyOrder), ctx, request)
}

// PlaceLimitSellOrder mocks base method.
func (m *MockOrderPlacer) PlaceLimitSellOrder(ctx context.Context, request tinkoffinvest.PlaceOrderRequest) (tinkoffinvest.OrderID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceLimitSellOrder", ctx, request)
	ret0, _ := ret[0].(tinkoffinvest.OrderID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlaceLimitSellOrder indicates an expected call of PlaceLimitSellOrder.
func (mr *MockOrderPlacerMockRecorder) PlaceLimitSellOrder(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceLimitSellOrder", reflect.TypeOf((*MockOrderPlacer)(nil).PlaceLimitSellOrder), ctx, request)
}

// SubscribeForOrderBookChanges mocks base method.
func (m *MockOrderPlacer) SubscribeForOrderBookChanges(ctx context.Context, reqs []tinkoffinvest.OrderBookRequest) (<-chan tinkoffinvest.OrderBookChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeForOrderBookChanges", ctx, reqs)
	ret0, _ := ret[0].(<-chan tinkoffinvest.OrderBookChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeForOrderBookChanges indicates an expected call of SubscribeForOrderBookChanges.
func (mr *MockOrderPlacerMockRecorder) SubscribeForOrderBookChanges(ctx, reqs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeForOrderBookChanges", reflect.TypeOf((*MockOrderPlacer)(nil).SubscribeForOrderBookChanges), ctx, reqs)
}

// MockToolsCache is a mock of ToolsCache interface.
type MockToolsCache struct {
	ctrl     *gomock.Controller
	recorder *MockToolsCacheMockRecorder
}

// MockToolsCacheMockRecorder is the mock recorder for MockToolsCache.
type MockToolsCacheMockRecorder struct {
	mock *MockToolsCache
}

// NewMockToolsCache creates a new mock instance.
func NewMockToolsCache(ctrl *gomock.Controller) *MockToolsCache {
	mock := &MockToolsCache{ctrl: ctrl}
	mock.recorder = &MockToolsCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolsCache) EXPECT() *MockToolsCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockToolsCache) Get(ctx context.Context, figi tinkoffinvest.FIGI) (toolscache.Tool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, figi)
	ret0, _ := ret[0].(toolscache.Tool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockToolsCacheMockRecorder) Get(ctx, figi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockToolsCache)(nil).Get), ctx, figi)
}
