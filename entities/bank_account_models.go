package entities

type BankAccount struct {
	Amount                    int64  `json:"-"`
	AccountIdenitificationNum string `json:"account_identif_num"`
	BankFullName              string `json:"bank_name"`
	BankIdentificationNum     string `json:"bank_identif_num"`
	IsBlocked                 string `json:"-"`
	IsFrozen                  string `json:"-"`
}

type Credit struct {
	bankProviderName          string
	accountIdenitificationNum string
	percent                   float64
	sumOfCredit               int
	requestId                 int
}

type Transfer struct {
	senderAccountNum   string
	sumOfTransfer      int
	receiverAccountNum string
}

type InstallmentPlan struct {
	info     Credit
	duration int
}

type PaymentRequest struct {
	amount   int
	clientId int
}
