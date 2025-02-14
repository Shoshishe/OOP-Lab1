package service

import (
	"main/app_interfaces"
)

type UserInfoService struct {
	app_interfaces.UserInfo
	repos app_interfaces.UserInfo
}

func NewUserInfoService(repos app_interfaces.UserInfo) *UserInfoService {
	return &UserInfoService{repos: repos}
}
