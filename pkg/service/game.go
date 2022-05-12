package service

import (
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"time"
)

type GameService struct {
	repo *repository.Repository
}

const (
	MaxRoomsInGame    = 10
	MinLengthEdge     = 30.0
	MaxLengthEdge     = 120.0
	MaxCountError     = 15
	PercentDispersion = 0.1
	CountTries        = 30
)

func (g *GameService) GenerateRoomsForGame(startUser1, startUser2, count,
	campusId int) ([]model.Room, []model.Room, error) {
	countErrors := 0
	for i := 0; i < CountTries; i++ {
		if countErrors > MaxCountError {
			return nil, nil, errors.New("can't generate rooms")
		}

		rooms1, err := g.GenerateRandomRooms(startUser1, count, campusId)
		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		distance1, err := g.GetDistanceBetweenRooms(startUser1, rooms1)

		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		rooms2, err := g.GenerateRoomsByDistance(startUser2, rooms1, distance1)

		if err != nil {
			logger.WarningLogger.Println(err)
			countErrors++
			continue
		}

		distance2, err := g.GetDistanceBetweenRooms(startUser2, rooms2)

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

func NewGameService(repo *repository.Repository) *GameService {
	rand.Seed(time.Now().Unix())
	return &GameService{repo: repo}
}
