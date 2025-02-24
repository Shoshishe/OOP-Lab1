package serviceInterfaces

import "main/domain/entities"

type Bank interface {
	GetBanksList(pagination int, userRole entities.UserRole) ([]entities.Bank, error)
	AddBank(bank entities.Bank, userRole entities.UserRole) error
}
