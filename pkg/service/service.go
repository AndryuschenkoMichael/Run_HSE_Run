package service

import "Run_Hse_Run/pkg/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
