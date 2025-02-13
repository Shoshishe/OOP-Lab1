package entities

type BankAccount struct {
	accountIdenitificationNum string
	bankFullName              string
	bankIdentificationNum     string
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
