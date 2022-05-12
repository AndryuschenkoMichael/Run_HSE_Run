package service

import (
	"Run_Hse_Run/pkg/model"
	"fmt"
	"sync"
)

type WebsocketService struct {
	sync.Mutex
	gameEntries map[int]*entry
	requests    chan *entry
	games       chan model.Game
}

func (w *WebsocketService) SendGame(game model.Game) error {
	return nil
}

func (w *WebsocketService) AddUser(userId, roomId int) {
	w.Lock()
	e := w.gameEntries[userId]
	if e != nil && !e.canceled() {
		e.cancel()
	}

	e = newEntry(userId, roomId)
	w.Unlock()

	w.requests <- e
}

func (w *WebsocketService) Cancel(userId int) {
	w.Lock()
	e := w.gameEntries[userId]
	if e != nil && !e.canceled() {
		e.cancel()
	}
	w.Unlock()
}

func (w *WebsocketService) Start() {
	var previous *entry
	for value := range w.requests {
		if !value.canceled() {
			if previous != nil && !previous.canceled() {
				w.games <- model.Game{
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

func (w *WebsocketService) Run() {
	for value := range w.games {
		fmt.Println(value)
	}
}

func NewWebsocketService() *WebsocketService {
	ws := &WebsocketService{
		gameEntries: make(map[int]*entry),
		requests:    make(chan *entry, 20),
		games:       make(chan model.Game, 20),
	}

	go ws.Start()
	go ws.Run()

	return ws
}
