package service

import (
	"Run_Hse_Run/pkg/mailer"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/queue"
	"Run_Hse_Run/pkg/repository"
	"Run_Hse_Run/pkg/websocket"
	"net/http"
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
	ChangeProfileImage(userId, image int) error
}

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

type Service struct {
	Sender
	Authorization
	Friends
	Users
	Game
}

func NewService(repo *repository.Repository, sender *mailer.Mailer,
	queue *queue.Queue, websocket *websocket.Server) *Service {
	return &Service{
		Sender:        NewSenderService(sender),
		Authorization: NewAuthService(repo),
		Friends:       NewFriendsService(repo),
		Users:         NewUsersService(repo),
		Game:          NewGameService(repo, queue, websocket),
	}
}
