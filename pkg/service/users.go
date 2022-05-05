package service

import (
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
)

type UsersService struct {
	repo *repository.Repository
}

func (u *UsersService) GetUsersByNicknamePattern(nickname string) ([]model.User, error) {
	return u.repo.GetUsersByNicknamePattern(nickname)
}

func (u *UsersService) GetUserById(userId int) (model.User, error) {
	return u.repo.GetUserById(userId)
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}
