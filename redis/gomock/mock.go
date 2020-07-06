// Code generated by MockGen. DO NOT EDIT.
// Source: redis/redis.go

// Package gomock is a generated GoMock package.
package gomock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockCache) Set(key, value string, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockCacheMockRecorder) Set(key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), key, value, expire)
}

// SetNX mocks base method
func (m *MockCache) SetNX(key, value string, expire time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", key, value, expire)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX
func (mr *MockCacheMockRecorder) SetNX(key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockCache)(nil).SetNX), key, value, expire)
}

// Get mocks base method
func (m *MockCache) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCacheMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), key)
}

// Del mocks base method
func (m *MockCache) Del(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockCacheMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockCache)(nil).Del), key)
}

// MockBasic is a mock of Basic interface
type MockBasic struct {
	ctrl     *gomock.Controller
	recorder *MockBasicMockRecorder
}

// MockBasicMockRecorder is the mock recorder for MockBasic
type MockBasicMockRecorder struct {
	mock *MockBasic
}

// NewMockBasic creates a new mock instance
func NewMockBasic(ctrl *gomock.Controller) *MockBasic {
	mock := &MockBasic{ctrl: ctrl}
	mock.recorder = &MockBasicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBasic) EXPECT() *MockBasicMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockBasic) Set(key, value string, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockBasicMockRecorder) Set(key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockBasic)(nil).Set), key, value, expire)
}

// SetNX mocks base method
func (m *MockBasic) SetNX(key, value string, expire time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", key, value, expire)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX
func (mr *MockBasicMockRecorder) SetNX(key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockBasic)(nil).SetNX), key, value, expire)
}

// Get mocks base method
func (m *MockBasic) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBasicMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBasic)(nil).Get), key)
}

// Del mocks base method
func (m *MockBasic) Del(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockBasicMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockBasic)(nil).Del), key)
}
