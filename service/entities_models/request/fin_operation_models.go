package request

import (
	"time"

	"github.com/shopspring/decimal"
)

type LoanModel struct {
	BankProviderName  string          `json:"bank_name"`
	AccountIdentifNum string          `json:"account_identif_num"`
	Rate              decimal.Decimal `json:"rate"`
	LoanAmount        int64           `json:"amount"`
	EndOfLoanTerm     time.Time       `json:"end_of_term"`
}
