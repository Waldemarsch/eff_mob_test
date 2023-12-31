// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	main_models "eff_mob_test/models"
	repository "eff_mob_test/pkg/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
)

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// CreateSingleUserRedis mocks base method.
func (m *MockUserStorage) CreateSingleUserRedis(user *main_models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSingleUserRedis", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSingleUserRedis indicates an expected call of CreateSingleUserRedis.
func (mr *MockUserStorageMockRecorder) CreateSingleUserRedis(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSingleUserRedis", reflect.TypeOf((*MockUserStorage)(nil).CreateSingleUserRedis), user)
}

// CreateUser mocks base method.
func (m *MockUserStorage) CreateUser(arg0 main_models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserStorageMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserStorage)(nil).CreateUser), arg0)
}

// DeleteSingleUserRedis mocks base method.
func (m *MockUserStorage) DeleteSingleUserRedis(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSingleUserRedis", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSingleUserRedis indicates an expected call of DeleteSingleUserRedis.
func (mr *MockUserStorageMockRecorder) DeleteSingleUserRedis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSingleUserRedis", reflect.TypeOf((*MockUserStorage)(nil).DeleteSingleUserRedis), arg0)
}

// DeleteUser mocks base method.
func (m *MockUserStorage) DeleteUser(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserStorageMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserStorage)(nil).DeleteUser), arg0)
}

// GetRedisConfig mocks base method.
func (m *MockUserStorage) GetRedisConfig() *redis.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedisConfig")
	ret0, _ := ret[0].(*redis.Options)
	return ret0
}

// GetRedisConfig indicates an expected call of GetRedisConfig.
func (mr *MockUserStorageMockRecorder) GetRedisConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedisConfig", reflect.TypeOf((*MockUserStorage)(nil).GetRedisConfig))
}

// GetSingleUser mocks base method.
func (m *MockUserStorage) GetSingleUser(arg0 int) (*main_models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSingleUser", arg0)
	ret0, _ := ret[0].(*main_models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSingleUser indicates an expected call of GetSingleUser.
func (mr *MockUserStorageMockRecorder) GetSingleUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleUser", reflect.TypeOf((*MockUserStorage)(nil).GetSingleUser), arg0)
}

// GetSingleUserRedis mocks base method.
func (m *MockUserStorage) GetSingleUserRedis(arg0 int) (*repository.JSONResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSingleUserRedis", arg0)
	ret0, _ := ret[0].(*repository.JSONResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSingleUserRedis indicates an expected call of GetSingleUserRedis.
func (mr *MockUserStorageMockRecorder) GetSingleUserRedis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleUserRedis", reflect.TypeOf((*MockUserStorage)(nil).GetSingleUserRedis), arg0)
}

// GetUser mocks base method.
func (m *MockUserStorage) GetUser(arg0, arg1 int, arg2 map[string]interface{}) []main_models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]main_models.User)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserStorageMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserStorage)(nil).GetUser), arg0, arg1, arg2)
}

// UpdateSingleUserRedis mocks base method.
func (m *MockUserStorage) UpdateSingleUserRedis(arg0 int, arg1 *main_models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSingleUserRedis", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSingleUserRedis indicates an expected call of UpdateSingleUserRedis.
func (mr *MockUserStorageMockRecorder) UpdateSingleUserRedis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSingleUserRedis", reflect.TypeOf((*MockUserStorage)(nil).UpdateSingleUserRedis), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockUserStorage) UpdateUser(arg0 main_models.User, arg1 []byte) (*main_models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*main_models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserStorageMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserStorage)(nil).UpdateUser), arg0, arg1)
}
