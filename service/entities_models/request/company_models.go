package request

type BankModel struct {
	LegalName        string `json:"name"`
	LegalAdress      string `json:"adress"`
	PayersAccountNum string `json:"account_identif_num"`
	CompanyType      string `json:"company_type"`
	BankIdentifNum   string `json:"bank_identif_num"`
	BankType         string `json:"bank_type"`
}
