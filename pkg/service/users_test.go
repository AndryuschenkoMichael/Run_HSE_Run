package service

import (
	"Run_Hse_Run/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsersService_ChangeProfileImage(t *testing.T) {
	mockRepo := NewMockUsersRepository(t)
	mockRepo.EXPECT().ChangeProfileImage(1, 2231).Return(nil).Once()

	svc := NewUsersService(mockRepo)

	err := svc.ChangeProfileImage(1, 2231)
	assert.NoError(t, err)
}

func TestUsersService_GetUserById(t *testing.T) {
	user1 := model.User{
		Id:       1,
		Nickname: "keker",
		Email:    "ru@ru.ru",
		Image:    234234,
	}
	mockRepo := NewMockUsersRepository(t)
	mockRepo.EXPECT().GetUserById(1).Return(user1, nil).Once()

	svc := NewUsersService(mockRepo)

	user, err := svc.GetUserById(1)
	assert.NoError(t, err)
	assert.Equal(t, user1, user)
}

func TestUsersService_GetUsersByNicknamePattern(t *testing.T) {
	user1 := model.User{
		Id:       1,
		Nickname: "keker",
		Email:    "ru@ru.ru",
		Image:    234234,
	}
	user2 := model.User{
		Id:       2,
		Nickname: "kekLALALLALAOLO",
		Email:    "rukiki@ru.ru",
		Image:    324,
	}
	mockRepo := NewMockUsersRepository(t)
	mockRepo.EXPECT().GetUsersByNicknamePattern("ke").Return([]model.User{user1, user2}, nil).Once()

	svc := NewUsersService(mockRepo)

	users, err := svc.GetUsersByNicknamePattern("ke")
	assert.NoError(t, err)
	assert.Equal(t, []model.User{user1, user2}, users)
}

func TestUsersService_RenameUser(t *testing.T) {
	mockRepo := NewMockUsersRepository(t)
	mockRepo.EXPECT().RenameUser(1, "mem").Return(nil).Once()

	svc := NewUsersService(mockRepo)

	err := svc.RenameUser(1, "mem")
	assert.NoError(t, err)
}
