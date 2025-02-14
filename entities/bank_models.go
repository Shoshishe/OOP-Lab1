package entities

type BankType = int

const (
	EmissionBank   = 1
	CommercialBank = 2
)

type Bank struct {
	Info Company  `json:"bank_info" binding:"required"`
	Type BankType `json:"bank_type" binding:"required"`
}
