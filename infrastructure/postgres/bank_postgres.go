package postgres

import (
	"database/sql"
	"fmt"
	"main/entities"
	"main/infrastructure"
)

type BankPostgres struct {
	infrastructure.Bank
	db *sql.DB
}

func (bankRepo *BankPostgres) AddBank(bank entities.Bank) error {
	query := fmt.Sprintf("INSERT INTO %s (name, adress, payers_acc_num, company_type, bank_ident_num, type) values ($1,$2,$3,$4,$5,$6)", BanksTable)
	_, err := bankRepo.db.Exec(query, bank.Info.LegalName, bank.Info.LegalAdress, bank.Info.PayersAccountNumber, bank.Info.Type, bank.Info.BankIdentificationNum, bank.Type)
		if err != nil {
		return err
	}
	return nil
}

func (bankRepo *BankPostgres) GetBanksList(pagination int) ([]entities.Bank, error) {
	offset := pagination * BanksPerRequestLimit
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %s OFFSET %s", BanksTable, fmt.Sprint(BanksPerRequestLimit), fmt.Sprint(offset))
	rows, err := bankRepo.db.Query(query)
	if err != nil {
		return nil, err
	}
	banksList := make([]entities.Bank, BanksPerRequestLimit)
	rows.Scan(banksList[0])
	i := 1
	for rows.Next() {
		rows.Scan(banksList[i])
		i++
	}
	return banksList, nil
}

func NewBankPostgres(db *sql.DB) *BankPostgres {
	return &BankPostgres{db: db}
}
