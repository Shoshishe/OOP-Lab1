package postgres

import (
	"database/sql"
	"fmt"
	"main/domain/entities"
	"main/service"
)

type BankPostgres struct {
	service.BankService
	db *sql.DB
}

func (bankRepo *BankPostgres) AddBank(bank entities.Bank) error {
	query := fmt.Sprintf("INSERT INTO %s (name, adress, payers_acc_num, company_type, bank_ident_num, type) values ($1,$2,$3,$4,$5,$6)", BanksTable)
	_, err := bankRepo.db.Exec(query, bank.Info.LegalName, bank.Info.LegalAdress, bank.Info.PayersAccountNumber, bank.Info.CompanyType, bank.Info.BankIdentificationNum, bank.Type)
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

func (bankRepo *BankPostgres) CheckBankExistance(bankIdentifNum entities.BankIdentificationNum) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT true FROM %s WHERE bank_ident_num=$1 limit 1", BanksTable)
	row := bankRepo.db.QueryRow(query, bankIdentifNum)
	err := row.Scan(exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (bankRepo *BankPostgres) CheckNameUniqueness(legalName string) (bool, error) {
	var count_of_rows int
	query := fmt.Sprintf("SELECT FROM %s WHERE name=$1 RETURNING ROW_NUMBER()", BanksTable)
	rows, err := bankRepo.db.Query(query,legalName)
	if err != nil {
		return false, err
	}
	err = rows.Scan(count_of_rows)
	if err != nil {
		return false, err
	}
	return count_of_rows == 1, nil
}

func NewBankPostgres(db *sql.DB) *BankPostgres {
	return &BankPostgres{db: db}
}
