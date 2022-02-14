// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/service.go

// Package mock_service is a generated GoMock package.
package mock

import (
	model "assignment/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWalletService is a mock of IWalletService interface.
type MockIWalletService struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletServiceMockRecorder
}

// MockIWalletServiceMockRecorder is the mock recorder for MockIWalletService.
type MockIWalletServiceMockRecorder struct {
	mock *MockIWalletService
}

// NewMockIWalletService creates a new mock instance.
func NewMockIWalletService(ctrl *gomock.Controller) *MockIWalletService {
	mock := &MockIWalletService{ctrl: ctrl}
	mock.recorder = &MockIWalletServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWalletService) EXPECT() *MockIWalletServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIWalletService) Create(username string) *model.Wallet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", username)
	ret0, _ := ret[0].(*model.Wallet)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIWalletServiceMockRecorder) Create(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIWalletService)(nil).Create), username)
}

// Get mocks base method.
func (m *MockIWalletService) Get(username string) (*model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", username)
	ret0, _ := ret[0].(*model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIWalletServiceMockRecorder) Get(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIWalletService)(nil).Get), username)
}

// GetAll mocks base method.
func (m *MockIWalletService) GetAll() []*model.Wallet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.Wallet)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIWalletServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIWalletService)(nil).GetAll))
}

// Update mocks base method.
func (m *MockIWalletService) Update(username string, balance int) (*model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", username, balance)
	ret0, _ := ret[0].(*model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIWalletServiceMockRecorder) Update(username, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIWalletService)(nil).Update), username, balance)
}