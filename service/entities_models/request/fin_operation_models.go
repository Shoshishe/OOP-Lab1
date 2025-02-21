package request

import (
	"time"

	"github.com/shopspring/decimal"
)

type LoanModel struct {
	BankProviderName  string          `json:"bank_name" binding:"required"`
	AccountIdentifNum string          `json:"account_identif_num" binding:"required"`
	Rate              decimal.Decimal `json:"rate" binding:"required"`
	LoanAmount        int64           `json:"amount" binding:"required"`
	EndOfLoanTerm     time.Time       `json:"end_of_term" binding:"required"`
}

type TransferModel struct {
	TransferOwnerId    int    `json:"-"`
	SenderAccountNum   string `json:"sender_acc_num" binding:"required"`
	SumOfTransfer      int64  `json:"amount" binding:"required"`
	ReceiverAccountNum string `json:"receiver_acc_num" binding:"required"`
}
