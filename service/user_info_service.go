package service

import (
	"main/app_interfaces"
	"main/entities"
)

type UserInfoService struct {
	app_interfaces.UserInfo
	repos app_interfaces.UserInfo
}

func (serv *UserInfoService) getRole(userId entities.UserRole) (entities.UserRole, error) {
	return serv.repos.GetRole(userId)
}

func NewUserInfoService(repos app_interfaces.UserInfo) *UserInfoService {
	return &UserInfoService{repos: repos}
}
