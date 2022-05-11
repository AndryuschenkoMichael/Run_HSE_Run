package service

import (
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

type GameService struct {
	repo *repository.Repository
}

func (g *GameService) GetRoomByCodePattern(code string) ([]model.Room, error) {
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

	return g.repo.GetRoomByCodePattern(code)
}

func NewGameService(repo *repository.Repository) *GameService {
	rand.Seed(time.Now().Unix())
	return &GameService{repo: repo}
}
