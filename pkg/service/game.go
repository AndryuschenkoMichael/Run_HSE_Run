package service

import (
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/queue"
	"Run_Hse_Run/pkg/repository"
	"Run_Hse_Run/pkg/websocket"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
	"time"
)

const (
	MaxRoomsInGame    = 10
	MinLengthEdge     = 30.0
	MaxLengthEdge     = 120.0
	MaxCountError     = 15
	PercentDispersion = 0.1
	CountTries        = 30
	TimeOut           = 3 * time.Minute
	TimeTick          = time.Second
	InfTime           = 1000000
)

var (
	mu           sync.Mutex
	isSendResult = make(map[int]struct{})
)

type GameService struct {
	repo      *repository.Repository
	queue     *queue.Queue
	websocket *websocket.Server
}

func (g *GameService) UpdateTime(gameId, userId, time int) error {
	return g.repo.UpdateTime(gameId, userId, time)
}

func (g *GameService) SendResult(gameId, userIdFirst, timeUser1 int) {
	game, err := g.repo.GetGame(gameId)
	if err != nil {
		return
	}

	userIdSecond := game.UserIdFirst
	if userIdSecond == userIdFirst {
		userIdSecond = game.UserIdSecond
	}

	mu.Lock()
	if _, ok := isSendResult[gameId]; ok {
		mu.Unlock()
		return
	}
	mu.Unlock()

	timer := time.NewTimer(TimeOut)
	ticker := time.NewTicker(TimeTick)

	for {
		select {
		case <-ticker.C:
			timeUser2, err := g.repo.GetTime(gameId, userIdSecond)
			if err != nil {
				logger.WarningLogger.Printf("can't get user time %s", err.Error())
				continue
			}

			mu.Lock()
			if _, ok := isSendResult[gameId]; ok {
				mu.Unlock()
				return
			}

			mu.Unlock()

			message1 := struct {
				GameResult string `json:"game_result"`
			}{}

			message2 := struct {
				GameResult string `json:"game_result"`
			}{}

			if timeUser1 == InfTime {
				message1.GameResult = "LOSE"
				message2.GameResult = "WIN"
			} else {
				if timeUser2.Time == -1 {
					continue
				}

				if timeUser1 == timeUser2.Time {
					message1.GameResult = "DRAW"
					message2.GameResult = "DRAW"
				} else if timeUser1 < timeUser2.Time {
					message1.GameResult = "WIN"
					message2.GameResult = "LOSE"
				} else {
					message1.GameResult = "LOSE"
					message2.GameResult = "WIN"
				}
			}

			mu.Lock()
			if _, ok := isSendResult[gameId]; ok {
				mu.Unlock()
				return
			}

			wg := sync.WaitGroup{}
			wg.Add(2)

			isSendResult[gameId] = struct{}{}
			mu.Unlock()

			go func() {
				g.websocket.WriteJson(userIdFirst, message1)
				wg.Done()
			}()

			go func() {
				g.websocket.WriteJson(userIdSecond, message2)
				wg.Done()
			}()

			wg.Wait()

			return
		case <-timer.C:
			message1 := struct {
				GameResult string `json:"game_result"`
			}{}

			message2 := struct {
				GameResult string `json:"game_result"`
			}{}

			if timeUser1 == InfTime {
				message1.GameResult = "DRAW"
				message2.GameResult = "DRAW"
			} else {
				message1.GameResult = "WIN"
				message2.GameResult = "LOSE"
			}

			mu.Lock()
			if _, ok := isSendResult[gameId]; ok {
				mu.Unlock()
				return
			}

			wg := sync.WaitGroup{}
			wg.Add(2)

			isSendResult[gameId] = struct{}{}
			mu.Unlock()

			go func() {
				g.websocket.WriteJson(userIdFirst, message1)
				wg.Done()
			}()

			go func() {
				g.websocket.WriteJson(userIdSecond, message2)
				wg.Done()
			}()

			wg.Wait()

			return
		}
	}
}

func (g *GameService) AddUser(userId, roomId int) {
	g.queue.AddUser(userId, roomId)
}

func (g *GameService) Cancel(userId int) {
	g.queue.Cancel(userId)
}

func (g *GameService) SendGame(game model.Game) error {
	if game.UserIdFirst == game.UserIdSecond {
		return errors.New("invalid game")
	}

	rooms1, rooms2, err := g.GenerateRoomsForGame(game.RoomIdFirst, game.RoomIdSecond, 3, 1)
	if err != nil {
		return err
	}

	id, err := g.repo.AddGame(game.UserIdFirst, game.UserIdSecond)
	if err != nil {
		return err
	}

	user1, err := g.repo.GetUserById(game.UserIdFirst)
	if err != nil {
		return err
	}

	user2, err := g.repo.GetUserById(game.UserIdSecond)
	if err != nil {
		return err
	}

	err = g.repo.AddTime(id, game.UserIdFirst, -1)
	if err != nil {
		return err
	}

	err = g.repo.AddTime(id, game.UserIdSecond, -1)
	if err != nil {
		return err
	}

	gameInfo1 := struct {
		Nickname string       `json:"nickname"`
		GameId   int          `json:"game_id"`
		Rooms    []model.Room `json:"rooms"`
	}{
		Nickname: user2.Nickname,
		GameId:   id,
		Rooms:    rooms1,
	}

	gameInfo2 := struct {
		Nickname string       `json:"nickname"`
		GameId   int          `json:"game_id"`
		Rooms    []model.Room `json:"rooms"`
	}{
		Nickname: user1.Nickname,
		GameId:   id,
		Rooms:    rooms2,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		g.websocket.WriteJson(game.UserIdFirst, gameInfo1)
		wg.Done()
	}()

	go func() {
		g.websocket.WriteJson(game.UserIdSecond, gameInfo2)
		wg.Done()
	}()

	wg.Wait()

	logger.WarningLogger.Printf("send game by users with id1 = %d, id2 = %d", game.UserIdFirst, game.UserIdSecond)

	return nil
}

func (g *GameService) UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	g.websocket.UpgradeConnection(w, r)
}

func (g *GameService) DeleteCall(userIdFirst, userIdSecond int) error {
	return g.repo.DeleteCall(userIdFirst, userIdSecond)
}

func (g *GameService) AddCall(userIdFirst, userIdSecond, roomIdFirst int) (model.Game, error) {
	return g.repo.AddCall(userIdFirst, userIdSecond, roomIdFirst)
}

func (g *GameService) GenerateRoomsForGame(startRoom1, startRoom2, count,
	campusId int) ([]model.Room, []model.Room, error) {
	countErrors := 0
	for i := 0; i < CountTries; i++ {
		if countErrors > MaxCountError {
			return nil, nil, errors.New("can't generate rooms")
		}

		rooms1, err := g.GenerateRandomRooms(startRoom1, count, campusId)
		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		distance1, err := g.GetDistanceBetweenRooms(startRoom1, rooms1)

		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		rooms2, err := g.GenerateRoomsByDistance(startRoom2, rooms1, distance1)

		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		distance2, err := g.GetDistanceBetweenRooms(startRoom2, rooms2)

		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		if math.Abs(distance2-distance1) < math.Max(distance2, distance1)*PercentDispersion {
			return rooms1, rooms2, nil
		}
	}

	return nil, nil, errors.New("can't generate rooms")
}

func (g *GameService) GenerateRandomRooms(startRoomId, count, campusId int) ([]model.Room, error) {
	if MaxRoomsInGame < count {
		return nil, errors.New(fmt.Sprintf("max rooms in game must be less than %d", MaxRoomsInGame))
	}

	if count < 1 {
		return nil, errors.New("count must be not less than 1")
	}

	var generatedRooms []model.Room

	rooms, err := g.GetRoomByCodePattern("", campusId)

	if len(rooms) < count {
		return nil, errors.New(fmt.Sprintf("count must be less than count of Rooms %d", len(rooms)))
	}

	if err != nil {
		return nil, err
	}

	used := make(map[int]struct{})
	used[startRoomId] = struct{}{}
	previous := startRoomId

	for len(generatedRooms) < count {
		index := rand.Intn(len(rooms))
		if _, ok := used[rooms[index].Id]; !ok {
			if edge, err := g.repo.GetEdge(previous, rooms[index].Id); err == nil {
				if MinLengthEdge < edge.Cost && edge.Cost < MaxLengthEdge {
					generatedRooms = append(generatedRooms, rooms[index])
					used[rooms[index].Id] = struct{}{}
					previous = rooms[index].Id
				}
			} else {
				return nil, err
			}
		}
	}

	return generatedRooms, nil
}

func (g *GameService) GetDistanceBetweenRooms(startRoomId int, rooms []model.Room) (float64, error) {
	if len(rooms) < 1 {
		return 0, errors.New("rooms must be more than zero")
	}

	cost := 0.0
	previous := startRoomId

	for _, room := range rooms {
		edge, err := g.repo.GetEdge(previous, room.Id)
		if err != nil {
			return 0, err
		}
		cost += edge.Cost
		previous = room.Id
	}

	return cost, nil
}

func (g *GameService) GenerateRoomsByDistance(startRoomId int, rooms []model.Room,
	distance float64) ([]model.Room, error) {
	count := len(rooms)
	if count < 1 {
		return nil, errors.New("rooms must be more than zero")
	}

	var generatedRooms []model.Room
	used := make(map[int]struct{})

	for _, room := range rooms {
		used[room.Id] = struct{}{}
	}

	used[startRoomId] = struct{}{}
	previous := startRoomId

	for len(generatedRooms) < count {
		var availableEdges []model.Edge
		edges, err := g.repo.GetListOfEdges(previous)

		if err != nil {
			return nil, err
		}

		for _, edge := range edges {
			if _, ok := used[edge.EndRoomId]; !ok {
				availableEdges = append(availableEdges, edge)
			}
		}

		if len(availableEdges) == 0 {
			return nil, errors.New("can't build a route")
		}

		edge := g.getNearestEdge(availableEdges, distance/float64(count-len(generatedRooms)))
		previous := edge.EndRoomId
		used[previous] = struct{}{}
		distance -= edge.Cost
		room, err := g.repo.GetRoomById(previous)
		if err != nil {
			return nil, err
		}

		generatedRooms = append(generatedRooms, room)
	}

	return generatedRooms, nil
}

func (g *GameService) getNearestEdge(edges []model.Edge, distance float64) model.Edge {
	minEdge := edges[0]

	for _, edge := range edges {
		if math.Abs(edge.Cost-distance) < math.Abs(minEdge.Cost-distance) {
			minEdge = edge
		}
	}

	return minEdge
}

func (g *GameService) GetRoomByCodePattern(code string, campusId int) ([]model.Room, error) {
	if 15 < len(code) {
		return nil, nil
	}

	expr := fmt.Sprintf("^[a-zA-Z0-9]{%d}", len(code))
	validUser, err := regexp.Compile(expr)
	if err != nil {
		return nil, nil
	}

	if !validUser.MatchString(code) {
		return nil, nil
	}

	return g.repo.GetRoomByCodePattern(code, campusId)
}

func (g *GameService) run() {
	for value := range g.queue.GetGameChan() {
		err := g.SendGame(value)
		if err != nil {
			logger.WarningLogger.Printf("can't send game: %s", err.Error())
		}
	}
}

func NewGameService(repo *repository.Repository, queue *queue.Queue, websocket *websocket.Server) *GameService {
	rand.Seed(time.Now().Unix())

	gameService := GameService{
		repo:      repo,
		queue:     queue,
		websocket: websocket,
	}

	go gameService.queue.Start()
	go gameService.run()

	return &gameService
}
