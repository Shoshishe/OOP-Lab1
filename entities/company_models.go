package entities

type CompanyType = string
type Company struct {
	LegalName             string      `json:"name"`
	LegalAdress           string      `json:"adress"`
	PayersAccountNumber   string      `json:"payers_acc_num"`
	Type                  CompanyType `json:"company_type"`
	BankIdentificationNum string      `json:"bank_identif_num"`
}
