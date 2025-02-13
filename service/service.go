package service

import (
	"main/entities"
	"main/infrastructure"
)

type Authorization interface {
	AddForApproval(entities.User) (int, error)
	AddUser(entities.User) (int, error)
	GetUser(username, password string) (*entities.User, error)
}

type Service struct {
	Authorization
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
