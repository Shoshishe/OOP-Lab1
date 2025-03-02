package response

import (
	"github.com/shopspring/decimal"
	"time"
)

type LoanModel struct {
	LoanId            int             `json:"-" db:"loan_id"`
	BankProviderName  string          `json:"bank_name" db:"bank_name"`
	AccountIdentifNum string          `json:"account_identif_num" db:"account_identif_num"`
	Rate              decimal.Decimal `json:"rate" db:"rate"`
	LoanAmount        int64           `json:"amount" db:"amount"`
	EndOfLoanTerm     time.Time       `json:"end_of_term" db:"end_of_term"`
	IsAccepted        bool            `json:"is_accepted" db:"is_accepted"`
}

type InstallmentPlanModel struct {
	BankProviderName string    `json:"bank_name"`
	AmountForPayment int64     `json:"amount"`
	CountOfPayments  int16     `json:"count_of_payments"`
	StartOfTerm      time.Time `json:"start_of_term"`
	EndOfTerm        time.Time `json:"end_of_term"`
}
type TransferModel struct {
	SenderAccountNum   string `json:"sender_acc_num" binding:"required"`
	SumOfTransfer      int64  `json:"amount" binding:"required"`
	ReceiverAccountNum string `json:"receiver_acc_num" binding:"required"`
}
