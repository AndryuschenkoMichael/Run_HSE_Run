package service

import (
	"Run_Hse_Run/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFriendsService_AddFriend(t *testing.T) {
	mockRepo := NewMockFriendsRepository(t)
	mockRepo.EXPECT().AddFriend(1, 2).Return(nil).Once()

	svc := NewFriendsService(mockRepo)

	err := svc.AddFriend(1, 2)
	assert.NoError(t, err)
	err = svc.AddFriend(3, 3)
	assert.NoError(t, err)
}

func TestFriendsService_DeleteFriend(t *testing.T) {
	mockRepo := NewMockFriendsRepository(t)
	mockRepo.EXPECT().DeleteFriend(1, 2).Return(nil).Once()

	svc := NewFriendsService(mockRepo)

	err := svc.DeleteFriend(1, 2)
	assert.NoError(t, err)
	err = svc.DeleteFriend(3, 3)
	assert.NoError(t, err)
}

func TestFriendsService_GetFriends(t *testing.T) {
	user1 := model.User{
		Id:       0,
		Nickname: "kek",
		Email:    "dsf@ru.ru",
		Image:    0,
	}

	user2 := model.User{
		Id:       1,
		Nickname: "loo",
		Email:    "sdf@ru.ru",
		Image:    0,
	}
	mockRepo := NewMockFriendsRepository(t)
	mockRepo.EXPECT().GetFriends(3).Return([]model.User{user1, user2}, nil).Once()
	mockRepo.EXPECT().GetFriends(1).Return([]model.User{}, nil).Once()

	svc := NewFriendsService(mockRepo)

	users, err := svc.GetFriends(3)
	assert.NoError(t, err)
	assert.Equal(t, []model.User{user1, user2}, users)

	users, err = svc.GetFriends(1)
	assert.NoError(t, err)
	assert.Equal(t, []model.User{}, users)
}
