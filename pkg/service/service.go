package service

import (
	"Run_Hse_Run/pkg/mailer"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
	"sync"
)

var (
	Mu    sync.Mutex
	Codes = make(map[string]int)
)

type Sender interface {
	SendEmail(email string) error
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
	GenerateToken(email string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Friends interface {
	AddFriend(userIdFrom, userIdTo int) error
	DeleteFriend(userIdFrom, userIdTo int) error
	GetFriends(userId int) ([]model.User, error)
}

type Users interface {
	GetUserById(userId int) (model.User, error)
	GetUsersByNicknamePattern(nickname string) ([]model.User, error)
	RenameUser(userId int, nickname string) error
}

type Service struct {
	Sender
	Authorization
	Friends
	Users
}

func NewService(repo *repository.Repository, sender *mailer.Mailer) *Service {
	return &Service{
		Sender:        NewSenderService(sender),
		Authorization: NewAuthService(repo),
		Friends:       NewFriendsService(repo),
		Users:         NewUsersService(repo),
	}
}
