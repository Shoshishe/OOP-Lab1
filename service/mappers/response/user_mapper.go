package response_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/response"
)

func ToUserAuthModel(usr entities.User) *response.UserAuthModel {
	return &response.UserAuthModel{
		FullName: usr.FullName(),
		PasportSeries: usr.PasportSeries(),
		PasportNum: usr.PasportNum(),
		MobilePhone: usr.MobilePhone(),
		Email: usr.Email(),
	}
}