package serviceInterfaces

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	"main/service/entities_models/response"
)

type Bank interface {
	GetBanksList(pagination int, userRole entities.UserRole) ([]response.BankModel, error)
	AddBank(bank request.BankModel,userId int, userRole entities.UserRole) error
}
