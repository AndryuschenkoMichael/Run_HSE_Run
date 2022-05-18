package repository

import (
	"Run_Hse_Run/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
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
	ChangeProfileImage(userId, image int) error
}

type Game interface {
	GetRoomByCodePattern(code string, campusId int) ([]model.Room, error)
	GetEdge(startRoomId, endRoomId int) (model.Edge, error)
	GetListOfEdges(startRoomId int) ([]model.Edge, error)
	GetRoomById(roomId int) (model.Room, error)
	AddCall(userIdFirst, userIdSecond, roomIdFirst int) (model.Game, error)
	DeleteCall(userIdFirst, userIdSecond int) error
	GetGame(gameId int) (model.GameUsers, error)
	GetTime(gameId, userId int) (model.Time, error)
	AddGame(userIdFirst, userIdSecond int) (int, error)
	AddTime(gameId, userId, time int) error
	UpdateTime(gameId, userId, time int) error
}

type Repository struct {
	Authorization
	Friends
	Users
	Game
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Friends:       NewFriendPostgres(db),
		Users:         NewUsersPostgres(db),
		Game:          NewGamePostgres(db),
	}
}
