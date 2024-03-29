// Code generated by mockery v2.20.0. DO NOT EDIT.

package service

import (
	model "Run_Hse_Run/pkg/model"

	mock "github.com/stretchr/testify/mock"
)

// MockGameRepository is an autogenerated mock type for the GameRepository type
type MockGameRepository struct {
	mock.Mock
}

type MockGameRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGameRepository) EXPECT() *MockGameRepository_Expecter {
	return &MockGameRepository_Expecter{mock: &_m.Mock}
}

// AddCall provides a mock function with given fields: userIdFirst, userIdSecond, roomIdFirst
func (_m *MockGameRepository) AddCall(userIdFirst int, userIdSecond int, roomIdFirst int) (model.Game, error) {
	ret := _m.Called(userIdFirst, userIdSecond, roomIdFirst)

	var r0 model.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, int) (model.Game, error)); ok {
		return rf(userIdFirst, userIdSecond, roomIdFirst)
	}
	if rf, ok := ret.Get(0).(func(int, int, int) model.Game); ok {
		r0 = rf(userIdFirst, userIdSecond, roomIdFirst)
	} else {
		r0 = ret.Get(0).(model.Game)
	}

	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(userIdFirst, userIdSecond, roomIdFirst)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_AddCall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddCall'
type MockGameRepository_AddCall_Call struct {
	*mock.Call
}

// AddCall is a helper method to define mock.On call
//   - userIdFirst int
//   - userIdSecond int
//   - roomIdFirst int
func (_e *MockGameRepository_Expecter) AddCall(userIdFirst interface{}, userIdSecond interface{}, roomIdFirst interface{}) *MockGameRepository_AddCall_Call {
	return &MockGameRepository_AddCall_Call{Call: _e.mock.On("AddCall", userIdFirst, userIdSecond, roomIdFirst)}
}

func (_c *MockGameRepository_AddCall_Call) Run(run func(userIdFirst int, userIdSecond int, roomIdFirst int)) *MockGameRepository_AddCall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockGameRepository_AddCall_Call) Return(_a0 model.Game, _a1 error) *MockGameRepository_AddCall_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_AddCall_Call) RunAndReturn(run func(int, int, int) (model.Game, error)) *MockGameRepository_AddCall_Call {
	_c.Call.Return(run)
	return _c
}

// AddGame provides a mock function with given fields: userIdFirst, userIdSecond
func (_m *MockGameRepository) AddGame(userIdFirst int, userIdSecond int) (int, error) {
	ret := _m.Called(userIdFirst, userIdSecond)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (int, error)); ok {
		return rf(userIdFirst, userIdSecond)
	}
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(userIdFirst, userIdSecond)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userIdFirst, userIdSecond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_AddGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGame'
type MockGameRepository_AddGame_Call struct {
	*mock.Call
}

// AddGame is a helper method to define mock.On call
//   - userIdFirst int
//   - userIdSecond int
func (_e *MockGameRepository_Expecter) AddGame(userIdFirst interface{}, userIdSecond interface{}) *MockGameRepository_AddGame_Call {
	return &MockGameRepository_AddGame_Call{Call: _e.mock.On("AddGame", userIdFirst, userIdSecond)}
}

func (_c *MockGameRepository_AddGame_Call) Run(run func(userIdFirst int, userIdSecond int)) *MockGameRepository_AddGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockGameRepository_AddGame_Call) Return(_a0 int, _a1 error) *MockGameRepository_AddGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_AddGame_Call) RunAndReturn(run func(int, int) (int, error)) *MockGameRepository_AddGame_Call {
	_c.Call.Return(run)
	return _c
}

// AddTime provides a mock function with given fields: gameId, userId, time
func (_m *MockGameRepository) AddTime(gameId int, userId int, time int) error {
	ret := _m.Called(gameId, userId, time)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int) error); ok {
		r0 = rf(gameId, userId, time)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGameRepository_AddTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTime'
type MockGameRepository_AddTime_Call struct {
	*mock.Call
}

// AddTime is a helper method to define mock.On call
//   - gameId int
//   - userId int
//   - time int
func (_e *MockGameRepository_Expecter) AddTime(gameId interface{}, userId interface{}, time interface{}) *MockGameRepository_AddTime_Call {
	return &MockGameRepository_AddTime_Call{Call: _e.mock.On("AddTime", gameId, userId, time)}
}

func (_c *MockGameRepository_AddTime_Call) Run(run func(gameId int, userId int, time int)) *MockGameRepository_AddTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockGameRepository_AddTime_Call) Return(_a0 error) *MockGameRepository_AddTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGameRepository_AddTime_Call) RunAndReturn(run func(int, int, int) error) *MockGameRepository_AddTime_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCall provides a mock function with given fields: userIdFirst, userIdSecond
func (_m *MockGameRepository) DeleteCall(userIdFirst int, userIdSecond int) error {
	ret := _m.Called(userIdFirst, userIdSecond)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(userIdFirst, userIdSecond)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGameRepository_DeleteCall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCall'
type MockGameRepository_DeleteCall_Call struct {
	*mock.Call
}

// DeleteCall is a helper method to define mock.On call
//   - userIdFirst int
//   - userIdSecond int
func (_e *MockGameRepository_Expecter) DeleteCall(userIdFirst interface{}, userIdSecond interface{}) *MockGameRepository_DeleteCall_Call {
	return &MockGameRepository_DeleteCall_Call{Call: _e.mock.On("DeleteCall", userIdFirst, userIdSecond)}
}

func (_c *MockGameRepository_DeleteCall_Call) Run(run func(userIdFirst int, userIdSecond int)) *MockGameRepository_DeleteCall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockGameRepository_DeleteCall_Call) Return(_a0 error) *MockGameRepository_DeleteCall_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGameRepository_DeleteCall_Call) RunAndReturn(run func(int, int) error) *MockGameRepository_DeleteCall_Call {
	_c.Call.Return(run)
	return _c
}

// GetEdge provides a mock function with given fields: startRoomId, endRoomId
func (_m *MockGameRepository) GetEdge(startRoomId int, endRoomId int) (model.Edge, error) {
	ret := _m.Called(startRoomId, endRoomId)

	var r0 model.Edge
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (model.Edge, error)); ok {
		return rf(startRoomId, endRoomId)
	}
	if rf, ok := ret.Get(0).(func(int, int) model.Edge); ok {
		r0 = rf(startRoomId, endRoomId)
	} else {
		r0 = ret.Get(0).(model.Edge)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(startRoomId, endRoomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_GetEdge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEdge'
type MockGameRepository_GetEdge_Call struct {
	*mock.Call
}

// GetEdge is a helper method to define mock.On call
//   - startRoomId int
//   - endRoomId int
func (_e *MockGameRepository_Expecter) GetEdge(startRoomId interface{}, endRoomId interface{}) *MockGameRepository_GetEdge_Call {
	return &MockGameRepository_GetEdge_Call{Call: _e.mock.On("GetEdge", startRoomId, endRoomId)}
}

func (_c *MockGameRepository_GetEdge_Call) Run(run func(startRoomId int, endRoomId int)) *MockGameRepository_GetEdge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockGameRepository_GetEdge_Call) Return(_a0 model.Edge, _a1 error) *MockGameRepository_GetEdge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_GetEdge_Call) RunAndReturn(run func(int, int) (model.Edge, error)) *MockGameRepository_GetEdge_Call {
	_c.Call.Return(run)
	return _c
}

// GetGame provides a mock function with given fields: gameId
func (_m *MockGameRepository) GetGame(gameId int) (model.GameUsers, error) {
	ret := _m.Called(gameId)

	var r0 model.GameUsers
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.GameUsers, error)); ok {
		return rf(gameId)
	}
	if rf, ok := ret.Get(0).(func(int) model.GameUsers); ok {
		r0 = rf(gameId)
	} else {
		r0 = ret.Get(0).(model.GameUsers)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(gameId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_GetGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGame'
type MockGameRepository_GetGame_Call struct {
	*mock.Call
}

// GetGame is a helper method to define mock.On call
//   - gameId int
func (_e *MockGameRepository_Expecter) GetGame(gameId interface{}) *MockGameRepository_GetGame_Call {
	return &MockGameRepository_GetGame_Call{Call: _e.mock.On("GetGame", gameId)}
}

func (_c *MockGameRepository_GetGame_Call) Run(run func(gameId int)) *MockGameRepository_GetGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockGameRepository_GetGame_Call) Return(_a0 model.GameUsers, _a1 error) *MockGameRepository_GetGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_GetGame_Call) RunAndReturn(run func(int) (model.GameUsers, error)) *MockGameRepository_GetGame_Call {
	_c.Call.Return(run)
	return _c
}

// GetListOfEdges provides a mock function with given fields: startRoomId
func (_m *MockGameRepository) GetListOfEdges(startRoomId int) ([]model.Edge, error) {
	ret := _m.Called(startRoomId)

	var r0 []model.Edge
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]model.Edge, error)); ok {
		return rf(startRoomId)
	}
	if rf, ok := ret.Get(0).(func(int) []model.Edge); ok {
		r0 = rf(startRoomId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Edge)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(startRoomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_GetListOfEdges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListOfEdges'
type MockGameRepository_GetListOfEdges_Call struct {
	*mock.Call
}

// GetListOfEdges is a helper method to define mock.On call
//   - startRoomId int
func (_e *MockGameRepository_Expecter) GetListOfEdges(startRoomId interface{}) *MockGameRepository_GetListOfEdges_Call {
	return &MockGameRepository_GetListOfEdges_Call{Call: _e.mock.On("GetListOfEdges", startRoomId)}
}

func (_c *MockGameRepository_GetListOfEdges_Call) Run(run func(startRoomId int)) *MockGameRepository_GetListOfEdges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockGameRepository_GetListOfEdges_Call) Return(_a0 []model.Edge, _a1 error) *MockGameRepository_GetListOfEdges_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_GetListOfEdges_Call) RunAndReturn(run func(int) ([]model.Edge, error)) *MockGameRepository_GetListOfEdges_Call {
	_c.Call.Return(run)
	return _c
}

// GetRoomByCodePattern provides a mock function with given fields: code, campusId
func (_m *MockGameRepository) GetRoomByCodePattern(code string, campusId int) ([]model.Room, error) {
	ret := _m.Called(code, campusId)

	var r0 []model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]model.Room, error)); ok {
		return rf(code, campusId)
	}
	if rf, ok := ret.Get(0).(func(string, int) []model.Room); ok {
		r0 = rf(code, campusId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Room)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(code, campusId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_GetRoomByCodePattern_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRoomByCodePattern'
type MockGameRepository_GetRoomByCodePattern_Call struct {
	*mock.Call
}

// GetRoomByCodePattern is a helper method to define mock.On call
//   - code string
//   - campusId int
func (_e *MockGameRepository_Expecter) GetRoomByCodePattern(code interface{}, campusId interface{}) *MockGameRepository_GetRoomByCodePattern_Call {
	return &MockGameRepository_GetRoomByCodePattern_Call{Call: _e.mock.On("GetRoomByCodePattern", code, campusId)}
}

func (_c *MockGameRepository_GetRoomByCodePattern_Call) Run(run func(code string, campusId int)) *MockGameRepository_GetRoomByCodePattern_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockGameRepository_GetRoomByCodePattern_Call) Return(_a0 []model.Room, _a1 error) *MockGameRepository_GetRoomByCodePattern_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_GetRoomByCodePattern_Call) RunAndReturn(run func(string, int) ([]model.Room, error)) *MockGameRepository_GetRoomByCodePattern_Call {
	_c.Call.Return(run)
	return _c
}

// GetRoomById provides a mock function with given fields: roomId
func (_m *MockGameRepository) GetRoomById(roomId int) (model.Room, error) {
	ret := _m.Called(roomId)

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.Room, error)); ok {
		return rf(roomId)
	}
	if rf, ok := ret.Get(0).(func(int) model.Room); ok {
		r0 = rf(roomId)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(roomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_GetRoomById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRoomById'
type MockGameRepository_GetRoomById_Call struct {
	*mock.Call
}

// GetRoomById is a helper method to define mock.On call
//   - roomId int
func (_e *MockGameRepository_Expecter) GetRoomById(roomId interface{}) *MockGameRepository_GetRoomById_Call {
	return &MockGameRepository_GetRoomById_Call{Call: _e.mock.On("GetRoomById", roomId)}
}

func (_c *MockGameRepository_GetRoomById_Call) Run(run func(roomId int)) *MockGameRepository_GetRoomById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockGameRepository_GetRoomById_Call) Return(_a0 model.Room, _a1 error) *MockGameRepository_GetRoomById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_GetRoomById_Call) RunAndReturn(run func(int) (model.Room, error)) *MockGameRepository_GetRoomById_Call {
	_c.Call.Return(run)
	return _c
}

// GetTime provides a mock function with given fields: gameId, userId
func (_m *MockGameRepository) GetTime(gameId int, userId int) (model.Time, error) {
	ret := _m.Called(gameId, userId)

	var r0 model.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (model.Time, error)); ok {
		return rf(gameId, userId)
	}
	if rf, ok := ret.Get(0).(func(int, int) model.Time); ok {
		r0 = rf(gameId, userId)
	} else {
		r0 = ret.Get(0).(model.Time)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(gameId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGameRepository_GetTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTime'
type MockGameRepository_GetTime_Call struct {
	*mock.Call
}

// GetTime is a helper method to define mock.On call
//   - gameId int
//   - userId int
func (_e *MockGameRepository_Expecter) GetTime(gameId interface{}, userId interface{}) *MockGameRepository_GetTime_Call {
	return &MockGameRepository_GetTime_Call{Call: _e.mock.On("GetTime", gameId, userId)}
}

func (_c *MockGameRepository_GetTime_Call) Run(run func(gameId int, userId int)) *MockGameRepository_GetTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockGameRepository_GetTime_Call) Return(_a0 model.Time, _a1 error) *MockGameRepository_GetTime_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGameRepository_GetTime_Call) RunAndReturn(run func(int, int) (model.Time, error)) *MockGameRepository_GetTime_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTime provides a mock function with given fields: gameId, userId, time
func (_m *MockGameRepository) UpdateTime(gameId int, userId int, time int) error {
	ret := _m.Called(gameId, userId, time)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int) error); ok {
		r0 = rf(gameId, userId, time)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGameRepository_UpdateTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTime'
type MockGameRepository_UpdateTime_Call struct {
	*mock.Call
}

// UpdateTime is a helper method to define mock.On call
//   - gameId int
//   - userId int
//   - time int
func (_e *MockGameRepository_Expecter) UpdateTime(gameId interface{}, userId interface{}, time interface{}) *MockGameRepository_UpdateTime_Call {
	return &MockGameRepository_UpdateTime_Call{Call: _e.mock.On("UpdateTime", gameId, userId, time)}
}

func (_c *MockGameRepository_UpdateTime_Call) Run(run func(gameId int, userId int, time int)) *MockGameRepository_UpdateTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockGameRepository_UpdateTime_Call) Return(_a0 error) *MockGameRepository_UpdateTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGameRepository_UpdateTime_Call) RunAndReturn(run func(int, int, int) error) *MockGameRepository_UpdateTime_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockGameRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockGameRepository creates a new instance of MockGameRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockGameRepository(t mockConstructorTestingTNewMockGameRepository) *MockGameRepository {
	mock := &MockGameRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
