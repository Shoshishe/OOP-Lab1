package request

import "main/domain/entities"

type BankAccountModel struct {
	AccountIdentificationNum entities.AccountIdenitificationNum `json:"-"`
	BankFullName             entities.BankName                  `json:"bank_name" binding:"required"`
	BankIdentificationNum    entities.BankIdentificationNum     `json:"bank_identif_num" binding:"required"`
}

type CompanyBankAccountModel struct {
	AccountIdentificationNum entities.AccountIdenitificationNum `json:"-"`
	BankFullName             entities.BankName                  `json:"bank_name" binding:"required"`
	BankIdentificationNum    entities.BankIdentificationNum     `json:"bank_identif_num" binding:"required"`
	CompanyId                int                                `json:"company_id" binding:"required"`
}
