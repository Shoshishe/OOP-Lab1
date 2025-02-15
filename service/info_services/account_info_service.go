package infoService

import "main/app_interfaces"

type accountInfoService struct {
	app_interfaces.AccountInfo
	repos app_interfaces.AccountInfo
}

func (serv *accountInfoService) CheckBelonging(userId int, accountIdentifNum string) (bool, error) {
	return serv.repos.CheckBelonging(userId, accountIdentifNum)
}

func NewAccountInfoService(repos app_interfaces.AccountInfo) *accountInfoService {
	return &accountInfoService{repos: repos}
}
