package request_mappers

import (
	"crypto/sha512"
	"fmt"
	"main/service/entities_models/request"
	"main/domain/entities"
)

const saltCrypto = "35edtryuiojhgytfe3"

func ToUserEntitiy(input request.UserSignUpModel, outsideInterface entities.UserOutside) (*entities.User, error) {
	usr, err := entities.NewUser(outsideInterface,
		entities.WithEmail(input.Email),
		entities.WithFullName(input.FullName),
		entities.WithPassword(input.Password),
		entities.WithPhone(input.MobilePhone),
		entities.WithPasportSeries(input.PasportSeries),
		entities.WithPasportNum(input.PasportNum),
	)
	return usr, err
}

func GenerateHashedPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltCrypto)))
}