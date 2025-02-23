package entities

import (
	"errors"
	domainErrors "main/domain/entities/domain_errors"
)

type CompanyType = string
type Adress = string
type Name = string
type BankType = string
type PayersAccountNumber = string

const (
	EmissionBank               = "emission_bank"
	CommercialBank             = "commercial_bank"
	LimitedLiabilityCompany    = "LLC"
	IndividualEnterpreneur     = "IE"
	ClosedJointStockCompany    = "CLJC"
	AdditionalLiabilityCompany = "ALC"
)

type Bank struct {
	Info Company
	Type BankType
}

type CompanyOutside interface {
	CheckNameUniqueness(legalName Name) (bool, error)
	CheckBankExistance(bankIdentifNum BankIdentificationNum) (bool, error)
}

type Company struct {
	id                    int
	legalName             Name
	legalAdress           Adress
	payersAccountNumber   PayersAccountNumber
	companyType           CompanyType
	bankIdentificationNum BankIdentificationNum
	outsideInfo           CompanyOutside
}

func NewCompany(legalName Name, legalAdress Adress,
	payersAccountNum AccountIdenitificationNum, companyType CompanyType, bankIdentifNum BankIdentificationNum) (*Company, error) {
	companyValue := &Company{
		legalName:             legalName,
		legalAdress:           legalAdress,
		payersAccountNumber:   payersAccountNum,
		companyType:           companyType,
		bankIdentificationNum: bankIdentifNum,
	}
	err := companyValue.ValidateCompany()
	if err != nil {
		return nil, err
	}
	return companyValue, nil
}

func (company *Company) ValidateCompany() error {
	return errors.Join(
		company.ValidateLegalName(),
		company.ValidateBankIdentifNum(),
		company.ValidateCompanyType(),
	)
}

func (company *Company) ValidateLegalName() error {
	isUnique, err := company.outsideInfo.CheckNameUniqueness(company.legalName)
	if err != nil {
		return err
	}
	if !isUnique {
		return domainErrors.NewInvalidField("legal name is not unique")
	}
	return nil
}

func (company *Company) ValidateBankIdentifNum() error {
	exists, err := company.outsideInfo.CheckBankExistance(company.bankIdentificationNum)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.NewInvalidField("bank identif num doesn't belong to anyone")
	}
	return nil
}

func (company *Company) ValidateCompanyType() error {
	var err error
	switch company.companyType {
	case LimitedLiabilityCompany, IndividualEnterpreneur:
		break
	case ClosedJointStockCompany, AdditionalLiabilityCompany:
		break
	default:
		err = domainErrors.NewInvalidField("invalid company type")
	}
	return err
}

func (company *Company) Id() int {
	return company.id
}

func (company *Company) LegalName() Name {
	return company.legalName
}

func (company *Company) LegalAdress() Adress {
	return company.legalAdress
}

func (company *Company) PayersAccountNumber() AccountIdenitificationNum {
	return company.payersAccountNumber
}

func (company *Company) BankIdentificationNum() BankIdentificationNum {
	return company.bankIdentificationNum
}

func (company *Company) CompanyType() CompanyType {
	return company.companyType
}

func (bank *Bank) BankType() BankType {
	return bank.Type
}
