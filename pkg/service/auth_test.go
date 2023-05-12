package service

import (
	"Run_Hse_Run/pkg/model"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_CreateUser(t *testing.T) {
	user1 := model.User{
		Id:       1,
		Nickname: "Nick",
		Email:    "dsfk@",
		Image:    1,
	}

	user2 := model.User{
		Id:       2,
		Nickname: "Hello",
		Email:    "2@3",
		Image:    123,
	}

	user3 := model.User{
		Id:       3,
		Nickname: "OSDFKDSJFKLDSAJKJFKJADFADJFKADJKFA",
		Email:    "2@3",
		Image:    123,
	}
	mockRepo := NewMockAuthorizationRepository(t)
	mockRepo.EXPECT().CreateUser(user1).Return(0, nil).Once()
	mockRepo.EXPECT().CreateUser(user2).Return(1, nil).Once()

	authSvc := NewAuthService(mockRepo)

	userID, err := authSvc.CreateUser(user1)
	assert.Equal(t, 0, userID)
	assert.NoError(t, err)

	userID, err = authSvc.CreateUser(user2)
	assert.Equal(t, 1, userID)
	assert.NoError(t, err)

	userID, err = authSvc.CreateUser(user3)

	assert.Equal(t, 0, userID)
	assert.Equal(t, err, errors.New(NicknameError))
}

func TestAuthService_Token(t *testing.T) {
	user := model.User{
		Id:       1,
		Nickname: "supauser",
		Email:    "user@mail.ru",
	}

	mockRepo := NewMockAuthorizationRepository(t)
	mockRepo.EXPECT().GetUser("user@mail.ru").Return(user, nil).Once()
	authSvc := NewAuthService(mockRepo)

	token, err := authSvc.GenerateToken("user@mail.ru")
	assert.NoError(t, err)
	userID, err := authSvc.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, 1, userID)
}

func TestAuthService_GetUser(t *testing.T) {
	user1 := model.User{
		Id:       1,
		Nickname: "Nick",
		Email:    "dsfk@ru.ru",
		Image:    1,
	}

	user2 := model.User{
		Id:       2,
		Nickname: "Hello",
		Email:    "212@mail.ru",
		Image:    123,
	}

	mockRepo := NewMockAuthorizationRepository(t)
	mockRepo.EXPECT().GetUser("212@mail.ru").Return(user2, nil).Once()
	mockRepo.EXPECT().GetUser("dsfk@ru.ru").Return(user1, nil).Once()

	authSvc := NewAuthService(mockRepo)

	user, err := authSvc.GetUser("212@mail.ru")
	assert.Equal(t, user2, user)
	assert.NoError(t, err)

	user, err = authSvc.GetUser("dsfk@ru.ru")
	assert.Equal(t, user1, user)
	assert.NoError(t, err)
}
