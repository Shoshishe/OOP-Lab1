package request

type BankModel struct {
	LegalName        string `json:"name" binding:"required"`
	LegalAdress      string `json:"adress" binding:"required"`
	PayersAccountNum string `json:"payers_acc_num" binding:"required"`
	CompanyType      string `json:"company_type" binding:"required"`
	BankIdentifNum   string `json:"bank_identif_num" binding:"required"`
	BankType         string `json:"bank_type" binding:"required"`
}
