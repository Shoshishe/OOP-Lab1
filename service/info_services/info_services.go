package infoService 

import ("main/app_interfaces")

type InfoService struct {
	userInfoService
	accountInfoService
	repos app_interfaces.Info
}



func NewInfoService(repos app_interfaces.Info) *InfoService {
	return &InfoService{repos: repos}
}