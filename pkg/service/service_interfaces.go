package service

import (
	"Run_Hse_Run/pkg/model"
	"net/http"
)

//go:generate mockery --name Sender
type Sender interface {
	SendEmail(email string) error
}

//go:generate mockery --name Authorization
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
	GenerateToken(email string) (string, error)
	ParseToken(accessToken string) (int, error)
}

//go:generate mockery --name Friends
type Friends interface {
	AddFriend(userIdFrom, userIdTo int) error
	DeleteFriend(userIdFrom, userIdTo int) error
	GetFriends(userId int) ([]model.User, error)
}

//go:generate mockery --name Users
type Users interface {
	GetUserById(userId int) (model.User, error)
	GetUsersByNicknamePattern(nickname string) ([]model.User, error)
	RenameUser(userId int, nickname string) error
	ChangeProfileImage(userId, image int) error
}

//go:generate mockery --name Game
type Game interface {
	GetRoomByCodePattern(code string, campusId int) ([]model.Room, error)
	AddCall(userIdFirst, userIdSecond, roomIdFirst int) (model.Game, error)
	DeleteCall(userIdFirst, userIdSecond int) error
	AddUser(userId, roomId int)
	Cancel(userId int)
	SendGame(game model.Game) error
	UpgradeConnection(w http.ResponseWriter, r *http.Request)
	SendResult(gameId, userId, time int)
	UpdateTime(gameId, userId, time int) error
}
