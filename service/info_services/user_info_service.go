package infoService

import (
	"main/app_interfaces"
	"main/entities"
)

type userInfoService struct {
	app_interfaces.UserInfo
	repos app_interfaces.UserInfo
}

func (serv *userInfoService) getRole(userId entities.UserRole) (entities.UserRole, error) {
	return serv.repos.GetUserRole(userId)
}

func NewUserInfoService(repos app_interfaces.UserInfo) *userInfoService {
	return &userInfoService{repos: repos}
}
