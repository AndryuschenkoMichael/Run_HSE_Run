package service

import (
	"Run_Hse_Run/pkg/model"
	"net/http"
)

//go:generate mockery --name AuthorizationRepository
type AuthorizationRepository interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
}

//go:generate mockery --name FriendsRepository
type FriendsRepository interface {
	AddFriend(userIdFrom, userIdTo int) error
	DeleteFriend(userIdFrom, userIdTo int) error
	GetFriends(userId int) ([]model.User, error)
}

//go:generate mockery --name UsersRepository
type UsersRepository interface {
	GetUserById(userId int) (model.User, error)
	GetUsersByNicknamePattern(nickname string) ([]model.User, error)
	RenameUser(userId int, nickname string) error
	ChangeProfileImage(userId, image int) error
}

//go:generate mockery --name GameRepository
type GameRepository interface {
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

//go:generate mockery --name GameWebsocket
type GameWebsocket interface {
	WriteJson(userId int, message interface{})
	UpgradeConnection(w http.ResponseWriter, r *http.Request)
}

//go:generate mockery --name GameUsersService
type GameUsersService interface {
	GetUserById(userId int) (model.User, error)
}

//go:generate mockery --name SenderMailer
type SenderMailer interface {
	SendEmail(email string, text string) error
}
