package service

import (
	"Run_Hse_Run/pkg/model"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameService_AddCall(t *testing.T) {
	game1 := model.Game{
		UserIdFirst:  1,
		RoomIdFirst:  2,
		UserIdSecond: 3,
		RoomIdSecond: 4,
	}

	mockRepo := NewMockGameRepository(t)
	mockRepo.EXPECT().AddCall(1, 2, 3).Return(game1, nil).Once()

	svc := NewGameService(mockRepo, nil, nil)

	game, err := svc.AddCall(1, 2, 3)
	assert.NoError(t, err)
	assert.Equal(t, game1, game)
}

func TestGameService_DeleteCall(t *testing.T) {
	mockRepo := NewMockGameRepository(t)
	mockRepo.EXPECT().DeleteCall(1, 2).Return(nil).Once()

	svc := NewGameService(mockRepo, nil, nil)

	err := svc.DeleteCall(1, 2)
	assert.NoError(t, err)
}

func TestGameService_GenerateRandomRooms(t *testing.T) {
	room1 := model.Room{
		Id:       1,
		Code:     "R33",
		CampusId: 2,
	}
	room2 := model.Room{
		Id:       3,
		Code:     "R32",
		CampusId: 2,
	}

	edge1 := model.Edge{
		Id:          0,
		StartRoomId: 1,
		EndRoomId:   3,
		Cost:        54,
		CampusId:    2,
	}
	mockRepo := NewMockGameRepository(t)
	mockRepo.EXPECT().GetRoomByCodePattern("", 2).Return([]model.Room{room1, room2}, nil)
	mockRepo.EXPECT().GetEdge(1, 3).Return(edge1, nil)

	svc := NewGameService(mockRepo, nil, nil)

	rooms, err := svc.GenerateRandomRooms(1, 1, 2)
	fmt.Println(rooms)
	assert.NoError(t, err)
}

func TestGameService_GetRoomByCodePattern(t *testing.T) {
	room1 := model.Room{
		Id:       1,
		Code:     "R2",
		CampusId: 1,
	}
	room2 := model.Room{
		Id:       3,
		Code:     "R232",
		CampusId: 1,
	}

	mockRepo := NewMockGameRepository(t)
	mockRepo.EXPECT().GetRoomByCodePattern("R", 1).Return([]model.Room{room1, room2}, nil).Once()

	svc := NewGameService(mockRepo, nil, nil)

	game, err := svc.GetRoomByCodePattern("R", 1)
	assert.NoError(t, err)
	assert.Equal(t, []model.Room{room1, room2}, game)
}
