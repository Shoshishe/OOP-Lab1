package response

import (
	"github.com/shopspring/decimal"
	"time"
)

type LoanModel struct {
	BankProviderName  string          `json:"bank_name" db:"bank_name"`
	AccountIdentifNum string          `json:"account_identif_num" db:"account_identif_num"`
	Rate              decimal.Decimal `json:"rate" db:"rate"`
	LoanAmount        int64           `json:"amount" db:"amount"`
	EndOfLoanTerm     time.Time       `json:"end_of_term" db:"end_of_term"`
}
