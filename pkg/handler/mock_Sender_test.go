// Code generated by mockery v2.20.0. DO NOT EDIT.

package handler

import mock "github.com/stretchr/testify/mock"

// MockSender is an autogenerated mock type for the Sender type
type MockSender struct {
	mock.Mock
}

type MockSender_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSender) EXPECT() *MockSender_Expecter {
	return &MockSender_Expecter{mock: &_m.Mock}
}

// SendEmail provides a mock function with given fields: email
func (_m *MockSender) SendEmail(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSender_SendEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEmail'
type MockSender_SendEmail_Call struct {
	*mock.Call
}

// SendEmail is a helper method to define mock.On call
//   - email string
func (_e *MockSender_Expecter) SendEmail(email interface{}) *MockSender_SendEmail_Call {
	return &MockSender_SendEmail_Call{Call: _e.mock.On("SendEmail", email)}
}

func (_c *MockSender_SendEmail_Call) Run(run func(email string)) *MockSender_SendEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSender_SendEmail_Call) Return(_a0 error) *MockSender_SendEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSender_SendEmail_Call) RunAndReturn(run func(string) error) *MockSender_SendEmail_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockSender interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSender creates a new instance of MockSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSender(t mockConstructorTestingTNewMockSender) *MockSender {
	mock := &MockSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
