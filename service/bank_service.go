package service

import (
	"main/entities"
	"main/infrastructure"
	"main/usecases"
	"errors"
)

const (
	LimitedLiabilityCompany    = "LLC"
	IndividualEnterpreneur     = "IE"
	ClosedJointStockCompany    = "CLJC"
	AdditionalLiabilityCompany = "ALC"
)
type BankService struct {
	usecases.Bank
	repos infrastructure.Bank
}

func (serv *BankService) GetBanksList(userRole entities.UserRole, pagination int) ([]entities.Bank, error) {
	if userRole != entities.RoleAdmin {
		return nil, errors.New("not accessible")
	}
	return serv.repos.GetBanksList(pagination)
}

func (serv *BankService) AddBank(userRole entities.UserRole, bank entities.Bank) error {
	var err error
	switch bank.Info.Type {
	case LimitedLiabilityCompany, IndividualEnterpreneur:
		break
	case ClosedJointStockCompany, AdditionalLiabilityCompany:
		break
	default:
		err = errors.New("incorrect company type")
	}
	if err != nil {
		return err
	}
	return serv.repos.AddBank(bank)
}

func NewBankService(repos infrastructure.Bank) *BankService {
	return &BankService{repos: repos}
}
