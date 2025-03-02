package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	persistance "main/repositories/postgres/entities_models"
	persistanceMappers "main/repositories/postgres/mappers"
	"main/service/repository"
)

type BankPostgres struct {
	repository.BankRepository
	db *sql.DB
}

func (bankRepo *BankPostgres) AddBank(bank entities.Bank, usrId int) error {
	tx, err := bankRepo.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := fmt.Sprintf("INSERT INTO %s (name, adress, payers_acc_num, company_type, bank_ident_num, bank_type) values ($1,$2,$3,$4,$5,$6)", BanksTable)
	_, err = tx.Exec(query, bank.Info.LegalName(), bank.Info.LegalAdress(), bank.Info.PayersAccountNumber(), bank.Info.CompanyType(), bank.Info.BankIdentificationNum(), bank.Type)
	if err != nil {
		return err
	}
	args := make([]string, 0, 6)
	args = append(args, bank.Info.LegalName(), bank.Info.LegalAdress(), bank.Info.PayersAccountNumber(), bank.Info.CompanyType(), bank.Info.BankIdentificationNum(), bank.Type)
	err = InsertAction(tx, "AddBank", args, usrId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (bankRepo *BankPostgres) ReverseBankAddition(bank entities.Bank, usrId int) error {
	tx, err := bankRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE name=$1", BanksTable)
	_, err = tx.Exec(query, bank.Info.LegalName())
	if err != nil {
		return err
	}
	args := make([]string, 0, 6)
	args = append(args, bank.Info.LegalName(), bank.Info.LegalAdress(), bank.Info.PayersAccountNumber(), bank.Info.CompanyType(), bank.Info.CompanyType(), bank.Type)
	err = ReverseAction(tx, bankRepo.db, args, usrId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (bankRepo *BankPostgres) GetBanksList(pagination int) ([]entities.Bank, error) {
	var banksList []entities.Bank
	offset := (pagination - 1) * BanksPerRequestLimit
	query := fmt.Sprintf("SELECT id,name,adress,payers_acc_num,company_type,bank_ident_num,bank_type FROM %s WHERE bank_type != '' LIMIT %s OFFSET $1", BanksTable, fmt.Sprint(BanksPerRequestLimit))
	rows, err := bankRepo.db.Query(query, fmt.Sprint(offset))
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.New("no companies to show :(")
	}
	bank := persistance.BankPersistance{}
	err = rows.Scan(&bank.Id, &bank.LegalName, &bank.LegalAdress,
		&bank.PayersAccNumber, &bank.CompanyType, &bank.BankIdentifNum,
		&bank.BankType)
	if err != nil {
		return nil, err
	}
	tempBank, err := persistanceMappers.ToBankEntity(&bank)
	if err != nil {
		return nil, err
	}
	banksList = append(banksList, *tempBank)
	i := 1
	for rows.Next() {
		err = rows.Scan(&bank.Id, &bank.LegalName, &bank.LegalAdress,
			&bank.PayersAccNumber, &bank.CompanyType, &bank.BankIdentifNum,
			&bank.BankType)
		if err != nil {
			return nil, err
		}
		tempBank, err := persistanceMappers.ToBankEntity(&bank)
		banksList = append(banksList, *tempBank)
		if err != nil {
			return nil, err
		}
		i++
	}
	return banksList, nil
}

func (bankRepo *BankPostgres) CheckBankExistance(bankIdentifNum entities.BankIdentificationNum) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE bank_ident_num=$1 limit 1", BanksTable)
	if err := bankRepo.db.QueryRow(query, bankIdentifNum).Scan(&exists); err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return false, err
	}
}

func (bankRepo *BankPostgres) CheckNameUniqueness(legalName string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE name=$1", BanksTable)
	if err := bankRepo.db.QueryRow(query, legalName).Scan(&exists); err == nil {
		return false, nil
	} else if err == sql.ErrNoRows {
		return true, nil
	} else {
		return false, err
	}
}

func NewBankPostgres(db *sql.DB) *BankPostgres {
	return &BankPostgres{db: db}
}
