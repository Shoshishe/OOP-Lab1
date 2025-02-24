package postgres

import (
	"database/sql"
	"fmt"
	"main/domain/entities"
	persistance "main/repository/postgres/entities_models"
	persistanceMappers "main/repository/postgres/mappers"
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
	query := fmt.Sprintf("SELECT id,name,adress,payers_acc_num,company_type,bank_ident_num,bank_type FROM %s LIMIT %s OFFSET $1", BanksTable, fmt.Sprint(BanksPerRequestLimit))
	rows, err := bankRepo.db.Query(query, fmt.Sprint(offset))
	if err != nil {
		return nil, err
	}
	banksPersistanceList := make([]persistance.BankPersistance, BanksPerRequestLimit)
	banksList := make([]entities.Bank, BanksPerRequestLimit)
	err = rows.Scan(&banksPersistanceList[0].Id, &banksPersistanceList[0].LegalName, &banksPersistanceList[0].LegalAdress,
		&banksPersistanceList[0].PayersAccNumber, &banksPersistanceList[0].CompanyType, &banksPersistanceList[0].BankIdentifNum,
		&banksPersistanceList[0].BankType)
	if err != nil {
		return nil, err
	}
	tempBank, err := persistanceMappers.ToBankEntity(&banksPersistanceList[0])
	banksList[0] = *tempBank
	if err != nil {
		return nil, err
	}
	i := 1
	for rows.Next() {
		err = rows.Scan(&banksPersistanceList[i].Id, &banksPersistanceList[i].LegalName, &banksPersistanceList[i].LegalAdress,
			&banksPersistanceList[i].PayersAccNumber, &banksPersistanceList[i].CompanyType, &banksPersistanceList[i].BankIdentifNum,
			&banksPersistanceList[i].BankType)
		if err != nil {
			return nil, err
		}
		tempBank, err := persistanceMappers.ToBankEntity(&banksPersistanceList[0])
		banksList[0] = *tempBank
		if err != nil {
			return nil, err
		}
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
	rows, err := bankRepo.db.Query(query, legalName)
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
