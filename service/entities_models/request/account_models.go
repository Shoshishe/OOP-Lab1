package request

import "main/domain/entities"

type BankAccountModel struct {
	AccountIdentificationNum entities.AccountIdenitificationNum `json:"-"`
	BankFullName             entities.BankName `json:"bank_name"`
	BankIdentificationNum    entities.BankIdentificationNum `json:"bank_identif_num"`
}
