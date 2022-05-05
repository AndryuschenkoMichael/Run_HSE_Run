package service

import (
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
)

type FriendsService struct {
	repo *repository.Repository
}

func (f *FriendsService) AddFriend(userIdFrom, userIdTo int) error {
	if userIdFrom == userIdTo {
		return nil
	}

	return f.repo.AddFriend(userIdFrom, userIdTo)
}

func (f *FriendsService) DeleteFriend(userIdFrom, userIdTo int) error {
	if userIdFrom == userIdTo {
		return nil
	}

	return f.repo.DeleteFriend(userIdFrom, userIdTo)
}

func (f *FriendsService) GetFriends(userId int) ([]model.User, error) {
	return f.repo.GetFriends(userId)
}

func NewFriendsService(repo *repository.Repository) *FriendsService {
	return &FriendsService{repo: repo}
}
