package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

func ToUserEntitiy(input request.ClientSignUpModel, outsideInterface entities.UserOutside) (*entities.User, error) {
	usr, err := entities.NewUser(outsideInterface, input.Password, input.Email,
		entities.WithFullName(input.FullName),
		entities.WithPhone(input.MobilePhone),
		entities.WithPasportSeries(input.PasportSeries),
		entities.WithPasportNum(input.PasportNum),
	)
	return usr, err
}
