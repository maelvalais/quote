// Code generated by MockGen. DO NOT EDIT.
// Source: ../user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	memdb "github.com/hashicorp/go-memdb"
	service "github.com/maelvls/users-grpc/pkg/service"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserService) Create(arg0 *memdb.Txn, arg1 service.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), arg0, arg1)
}

// List mocks base method
func (m *MockUserService) List(arg0 *memdb.Txn) ([]service.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]service.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockUserServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserService)(nil).List), arg0)
}

// SearchAge mocks base method
func (m *MockUserService) SearchAge(txn *memdb.Txn, ageFrom, ageTo int32) ([]service.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAge", txn, ageFrom, ageTo)
	ret0, _ := ret[0].([]service.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAge indicates an expected call of SearchAge
func (mr *MockUserServiceMockRecorder) SearchAge(txn, ageFrom, ageTo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAge", reflect.TypeOf((*MockUserService)(nil).SearchAge), txn, ageFrom, ageTo)
}

// SearchName mocks base method
func (m *MockUserService) SearchName(txn *memdb.Txn, query string) ([]service.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchName", txn, query)
	ret0, _ := ret[0].([]service.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchName indicates an expected call of SearchName
func (mr *MockUserServiceMockRecorder) SearchName(txn, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchName", reflect.TypeOf((*MockUserService)(nil).SearchName), txn, query)
}

// GetByEmail mocks base method
func (m *MockUserService) GetByEmail(txn *memdb.Txn, email string) (service.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", txn, email)
	ret0, _ := ret[0].(service.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail
func (mr *MockUserServiceMockRecorder) GetByEmail(txn, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserService)(nil).GetByEmail), txn, email)
}
