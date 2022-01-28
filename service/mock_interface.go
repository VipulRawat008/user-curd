// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	entities "user-curd/entities"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUserService mocks base method.
func (m *MockUserService) CreateUserService(user entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserService", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserService indicates an expected call of CreateUserService.
func (mr *MockUserServiceMockRecorder) CreateUserService(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserService", reflect.TypeOf((*MockUserService)(nil).CreateUserService), user)
}

// DeleteUserService mocks base method.
func (m *MockUserService) DeleteUserService(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserService", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserService indicates an expected call of DeleteUserService.
func (mr *MockUserServiceMockRecorder) DeleteUserService(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserService", reflect.TypeOf((*MockUserService)(nil).DeleteUserService), id)
}

// GetAllUsersService mocks base method.
func (m *MockUserService) GetAllUsersService() ([]*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsersService")
	ret0, _ := ret[0].([]*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsersService indicates an expected call of GetAllUsersService.
func (mr *MockUserServiceMockRecorder) GetAllUsersService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersService", reflect.TypeOf((*MockUserService)(nil).GetAllUsersService))
}

// GetUserByIdService mocks base method.
func (m *MockUserService) GetUserByIdService(id int) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByIdService", id)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByIdService indicates an expected call of GetUserByIdService.
func (mr *MockUserServiceMockRecorder) GetUserByIdService(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByIdService", reflect.TypeOf((*MockUserService)(nil).GetUserByIdService), id)
}

// UpdateUserService mocks base method.
func (m *MockUserService) UpdateUserService(user entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserService", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserService indicates an expected call of UpdateUserService.
func (mr *MockUserServiceMockRecorder) UpdateUserService(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserService", reflect.TypeOf((*MockUserService)(nil).UpdateUserService), user)
}
