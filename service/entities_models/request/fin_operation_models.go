package request

import (
	"time"
)

type LoanModel struct {
	BankProviderName  string    `json:"bank_name" binding:"required"`
	AccountIdentifNum string    `json:"account_identif_num" binding:"required"`
	Rate              string    `json:"rate" binding:"required"`
	LoanAmount        int64     `json:"amount" binding:"required"`
	StartOfLoanTerm   time.Time `json:"start_of_term" binding:"required"`
	EndOfLoanTerm     time.Time `json:"end_of_term" binding:"required"`
}

type TransferModel struct {
	TransferOwnerId    int    `json:"-"`
	SenderAccountNum   string `json:"sender_acc_num" binding:"required"`
	SumOfTransfer      int64  `json:"amount" binding:"required"`
	ReceiverAccountNum string `json:"receiver_acc_num" binding:"required"`
}

type InstallmentPlanModel struct {
	BankProviderName  string    `json:"bank_name" binding:"required"`
	AmountForPayment  int64     `json:"amount" binding:"required"`
	CountOfPayments   int16     `json:"count_of_payments" binding:"required"`
	StartOfTerm       time.Time `json:"start_of_term" binding:"required"`
	AccountIdentifNum string    `json:"account_identif_num" binding:"required"`
	EndOfTerm         time.Time `json:"end_of_term" binding:"required"`
}

type PaymentRequestModel struct {
	Amount     int    `json:"money_amount" binding:"required"`
	AccountNum string `json:"account_identif_num" binding:"required"`
	FullName   string `json:"full_name" binding:"required"`
	ClientId   int    `json:"client_id" binding:"required"`
	CompanyId  int    `json:"company_id" binding:"required"`
}
