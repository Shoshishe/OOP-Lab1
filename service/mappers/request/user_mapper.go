package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	"main/utils"
)

func ToUserEntitiy(input request.UserSignUpModel, outsideInterface entities.UserOutside) (*entities.User, error) {
	usr, err := entities.NewUser(outsideInterface,
		entities.WithEmail(input.Email),
		entities.WithFullName(input.FullName),
		entities.WithPassword(utils.GenerateHashedPassword(input.Password)),
		entities.WithPhone(input.MobilePhone),
		entities.WithPasportSeries(input.PasportSeries),
		entities.WithPasportNum(input.PasportNum),
	)
	return usr, err
}