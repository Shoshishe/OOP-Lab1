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

func NewBank(info Company, Type BankType) (*Bank, error) {
	bank := &Bank{
		Info: info,
		Type: Type,
	}
	err := bank.Info.validator.ValidateCompany(&info)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func (bank *Bank) ValidateBankType() error {
	if bank.Type != CommercialBank && bank.Type != EmissionBank {
		return domainErrors.NewInvalidField("invalid bank type")
	}
	return nil
}

type CompanyOutside interface {
	CheckNameUniqueness(legalName Name) (bool, error)
	CheckBankExistance(bankIdentifNum BankIdentificationNum) (bool, error)
}

type BankOutside interface {
	CheckNameUniqueness(legalName Name) (bool, error)
}
type companyValidator interface {
	ValidateCompany(*Company) error
}

type bankValidatorPolicy struct {
	companyValidator
	validator *companyValidatorPolicy
}

func (bankValidator *bankValidatorPolicy) ValidateCompany(company *Company) error {
	return errors.Join(
		bankValidator.validator.ValidateLegalName(company),
		bankValidator.validator.ValidateCompanyType(company),
	)
}

func NewBankValidatorPolicy(validator *companyValidatorPolicy) *bankValidatorPolicy {
	return &bankValidatorPolicy{
		validator: validator,
	}
}

type companyValidatorPolicy struct {
	companyValidator
	outsideInfo CompanyOutside
}

func NewCompanyValidatorPolicy(outsideInfo CompanyOutside) *companyValidatorPolicy {
	return &companyValidatorPolicy{
		outsideInfo: outsideInfo,
	}
}

func (companyValidator *companyValidatorPolicy) ValidateCompany(company *Company) error {
	return errors.Join(
		companyValidator.ValidateLegalName(company),
		companyValidator.ValidateBankIdentifNum(company),
		companyValidator.ValidateCompanyType(company),
	)
}

func (companyValidator *companyValidatorPolicy) ValidateLegalName(company *Company) error {
	isUnique, err := companyValidator.outsideInfo.CheckNameUniqueness(company.legalName)
	if err != nil {
		return err
	}
	if !isUnique {
		return domainErrors.NewInvalidField("legal name is not unique")
	}
	return nil
}

func (companyValidator *companyValidatorPolicy) ValidateBankIdentifNum(company *Company) error {
	exists, err := companyValidator.outsideInfo.CheckBankExistance(company.bankIdentificationNum)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.NewInvalidField("bank identif num doesn't belong to anyone")
	}
	return nil
}

func (companyValidator *companyValidatorPolicy) ValidateCompanyType(company *Company) error {
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

func NewCompanyValidator(outsideInfo CompanyOutside) *companyValidatorPolicy {
	return &companyValidatorPolicy{
		outsideInfo: outsideInfo,
	}
}

type Company struct {
	id                    int
	legalName             Name
	legalAdress           Adress
	payersAccountNumber   PayersAccountNumber
	companyType           CompanyType
	bankIdentificationNum BankIdentificationNum
	validator             companyValidator
}

func NewCompany(legalName Name, legalAdress Adress,
	payersAccountNum AccountIdenitificationNum, companyType CompanyType, bankIdentifNum BankIdentificationNum, validator companyValidator) (*Company, error) {
	companyValue := &Company{
		validator:             validator,
		legalName:             legalName,
		legalAdress:           legalAdress,
		payersAccountNumber:   payersAccountNum,
		companyType:           companyType,
		bankIdentificationNum: bankIdentifNum,
	}
	err := validator.ValidateCompany(companyValue)
	if err != nil {
		return nil, err
	}
	return companyValue, nil
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
