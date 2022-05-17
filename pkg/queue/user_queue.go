package queue

import (
	"Run_Hse_Run/pkg/model"
	"sync"
)

type UserQueue struct {
	sync.Mutex
	gameEntries map[int]*entry
	requests    chan *entry
	games       chan model.Game
}

func (u *UserQueue) AddUser(userId, roomId int) {
	u.Lock()
	e := u.gameEntries[userId]
	if e != nil && !e.canceled() {
		e.cancel()
	}

	e = newEntry(userId, roomId)
	u.gameEntries[userId] = e
	u.Unlock()

	u.requests <- e
}

func (u *UserQueue) Cancel(userId int) {
	u.Lock()
	e := u.gameEntries[userId]
	if e != nil && !e.canceled() {
		e.cancel()
	}
	u.Unlock()
}

func (u *UserQueue) Start() {
	var previous *entry
	for value := range u.requests {
		if !value.canceled() {
			if previous != nil && !previous.canceled() {
				u.games <- model.Game{
					UserIdFirst:  previous.userId,
					RoomIdFirst:  previous.roomId,
					UserIdSecond: value.userId,
					RoomIdSecond: value.roomId,
				}
				previous = nil
			} else {
				previous = value
			}
		}
	}
}

func (u *UserQueue) GetGameChan() <-chan model.Game {
	return u.games
}

func NewUserQueue() *UserQueue {
	return &UserQueue{
		gameEntries: make(map[int]*entry),
		requests:    make(chan *entry),
		games:       make(chan model.Game),
	}
}
