// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/alazarbeyeneazu/Simple-Gym-management-system/ports (interfaces: DBPort)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	models "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockDBPort is a mock of DBPort interface.
type MockDBPort struct {
	ctrl     *gomock.Controller
	recorder *MockDBPortMockRecorder
}

// MockDBPortMockRecorder is the mock recorder for MockDBPort.
type MockDBPortMockRecorder struct {
	mock *MockDBPort
}

// NewMockDBPort creates a new mock instance.
func NewMockDBPort(ctrl *gomock.Controller) *MockDBPort {
	mock := &MockDBPort{ctrl: ctrl}
	mock.recorder = &MockDBPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBPort) EXPECT() *MockDBPortMockRecorder {
	return m.recorder
}

// CreatePymentType mocks base method.
func (m *MockDBPort) CreatePymentType(arg0 context.Context, arg1 models.PymentType) (models.PymentType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePymentType", arg0, arg1)
	ret0, _ := ret[0].(models.PymentType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePymentType indicates an expected call of CreatePymentType.
func (mr *MockDBPortMockRecorder) CreatePymentType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePymentType", reflect.TypeOf((*MockDBPort)(nil).CreatePymentType), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockDBPort) CreateUser(arg0 context.Context, arg1 models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDBPortMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDBPort)(nil).CreateUser), arg0, arg1)
}

// DeletePyment mocks base method.
func (m *MockDBPort) DeletePyment(arg0 context.Context, arg1 models.PymentType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePyment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePyment indicates an expected call of DeletePyment.
func (mr *MockDBPortMockRecorder) DeletePyment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePyment", reflect.TypeOf((*MockDBPort)(nil).DeletePyment), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockDBPort) DeleteUser(arg0 context.Context, arg1 models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockDBPortMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockDBPort)(nil).DeleteUser), arg0, arg1)
}

// GetAllPyments mocks base method.
func (m *MockDBPort) GetAllPyments(arg0 context.Context) ([]models.PymentType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPyments", arg0)
	ret0, _ := ret[0].([]models.PymentType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPyments indicates an expected call of GetAllPyments.
func (mr *MockDBPortMockRecorder) GetAllPyments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPyments", reflect.TypeOf((*MockDBPort)(nil).GetAllPyments), arg0)
}

// GetPymentById mocks base method.
func (m *MockDBPort) GetPymentById(arg0 context.Context, arg1 models.PymentType) (models.PymentType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPymentById", arg0, arg1)
	ret0, _ := ret[0].(models.PymentType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPymentById indicates an expected call of GetPymentById.
func (mr *MockDBPortMockRecorder) GetPymentById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPymentById", reflect.TypeOf((*MockDBPort)(nil).GetPymentById), arg0, arg1)
}

// GetUseByPhoneNumber mocks base method.
func (m *MockDBPort) GetUseByPhoneNumber(arg0 context.Context, arg1 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUseByPhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUseByPhoneNumber indicates an expected call of GetUseByPhoneNumber.
func (mr *MockDBPortMockRecorder) GetUseByPhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUseByPhoneNumber", reflect.TypeOf((*MockDBPort)(nil).GetUseByPhoneNumber), arg0, arg1)
}

// GetUserByFirstName mocks base method.
func (m *MockDBPort) GetUserByFirstName(arg0 context.Context, arg1 string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByFirstName", arg0, arg1)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByFirstName indicates an expected call of GetUserByFirstName.
func (mr *MockDBPortMockRecorder) GetUserByFirstName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByFirstName", reflect.TypeOf((*MockDBPort)(nil).GetUserByFirstName), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockDBPort) GetUserById(arg0 context.Context, arg1 uuid.UUID) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockDBPortMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockDBPort)(nil).GetUserById), arg0, arg1)
}

// GetUserByLastName mocks base method.
func (m *MockDBPort) GetUserByLastName(arg0 context.Context, arg1 string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLastName", arg0, arg1)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLastName indicates an expected call of GetUserByLastName.
func (mr *MockDBPortMockRecorder) GetUserByLastName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLastName", reflect.TypeOf((*MockDBPort)(nil).GetUserByLastName), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockDBPort) GetUsers(arg0 context.Context) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockDBPortMockRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockDBPort)(nil).GetUsers), arg0)
}

// UpdatePyment mocks base method.
func (m *MockDBPort) UpdatePyment(arg0 context.Context, arg1 models.PymentType) (models.PymentType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePyment", arg0, arg1)
	ret0, _ := ret[0].(models.PymentType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePyment indicates an expected call of UpdatePyment.
func (mr *MockDBPortMockRecorder) UpdatePyment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePyment", reflect.TypeOf((*MockDBPort)(nil).UpdatePyment), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockDBPort) UpdateUser(arg0 context.Context, arg1, arg2 models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockDBPortMockRecorder) UpdateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockDBPort)(nil).UpdateUser), arg0, arg1, arg2)
}
