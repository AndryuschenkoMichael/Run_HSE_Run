// Code generated by mockery v2.20.0. DO NOT EDIT.

package handler

import (
	model "Run_Hse_Run/pkg/model"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthorization is an autogenerated mock type for the Authorization type
type MockAuthorization struct {
	mock.Mock
}

type MockAuthorization_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthorization) EXPECT() *MockAuthorization_Expecter {
	return &MockAuthorization_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *MockAuthorization) CreateUser(user model.User) (int, error) {
	ret := _m.Called(user)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (int, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(model.User) int); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorization_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockAuthorization_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user model.User
func (_e *MockAuthorization_Expecter) CreateUser(user interface{}) *MockAuthorization_CreateUser_Call {
	return &MockAuthorization_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *MockAuthorization_CreateUser_Call) Run(run func(user model.User)) *MockAuthorization_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.User))
	})
	return _c
}

func (_c *MockAuthorization_CreateUser_Call) Return(_a0 int, _a1 error) *MockAuthorization_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorization_CreateUser_Call) RunAndReturn(run func(model.User) (int, error)) *MockAuthorization_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateToken provides a mock function with given fields: email
func (_m *MockAuthorization) GenerateToken(email string) (string, error) {
	ret := _m.Called(email)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorization_GenerateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateToken'
type MockAuthorization_GenerateToken_Call struct {
	*mock.Call
}

// GenerateToken is a helper method to define mock.On call
//   - email string
func (_e *MockAuthorization_Expecter) GenerateToken(email interface{}) *MockAuthorization_GenerateToken_Call {
	return &MockAuthorization_GenerateToken_Call{Call: _e.mock.On("GenerateToken", email)}
}

func (_c *MockAuthorization_GenerateToken_Call) Run(run func(email string)) *MockAuthorization_GenerateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAuthorization_GenerateToken_Call) Return(_a0 string, _a1 error) *MockAuthorization_GenerateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorization_GenerateToken_Call) RunAndReturn(run func(string) (string, error)) *MockAuthorization_GenerateToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: email
func (_m *MockAuthorization) GetUser(email string) (model.User, error) {
	ret := _m.Called(email)

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) model.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorization_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockAuthorization_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - email string
func (_e *MockAuthorization_Expecter) GetUser(email interface{}) *MockAuthorization_GetUser_Call {
	return &MockAuthorization_GetUser_Call{Call: _e.mock.On("GetUser", email)}
}

func (_c *MockAuthorization_GetUser_Call) Run(run func(email string)) *MockAuthorization_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAuthorization_GetUser_Call) Return(_a0 model.User, _a1 error) *MockAuthorization_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorization_GetUser_Call) RunAndReturn(run func(string) (model.User, error)) *MockAuthorization_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// ParseToken provides a mock function with given fields: accessToken
func (_m *MockAuthorization) ParseToken(accessToken string) (int, error) {
	ret := _m.Called(accessToken)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(accessToken)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(accessToken)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorization_ParseToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseToken'
type MockAuthorization_ParseToken_Call struct {
	*mock.Call
}

// ParseToken is a helper method to define mock.On call
//   - accessToken string
func (_e *MockAuthorization_Expecter) ParseToken(accessToken interface{}) *MockAuthorization_ParseToken_Call {
	return &MockAuthorization_ParseToken_Call{Call: _e.mock.On("ParseToken", accessToken)}
}

func (_c *MockAuthorization_ParseToken_Call) Run(run func(accessToken string)) *MockAuthorization_ParseToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAuthorization_ParseToken_Call) Return(_a0 int, _a1 error) *MockAuthorization_ParseToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorization_ParseToken_Call) RunAndReturn(run func(string) (int, error)) *MockAuthorization_ParseToken_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAuthorization interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAuthorization creates a new instance of MockAuthorization. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAuthorization(t mockConstructorTestingTNewMockAuthorization) *MockAuthorization {
	mock := &MockAuthorization{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}