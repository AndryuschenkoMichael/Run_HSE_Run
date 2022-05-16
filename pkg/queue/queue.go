package queue

import "Run_Hse_Run/pkg/model"

type Queuer interface {
	AddUser(userId, roomId int)
	Cancel(userId int)
	Start()
	GetGameChan() <-chan model.Game
}

type Queue struct {
	Queuer
}

func NewQueue() *Queue {
	return &Queue{
		Queuer: NewUserQueue(),
	}
}
