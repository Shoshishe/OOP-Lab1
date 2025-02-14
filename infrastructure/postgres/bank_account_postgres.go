package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/entities"
	"main/infrastructure"
)

type BankAccountPostgres struct {
	infrastructure.BankAccount
	db *sql.DB
}

func (bankAccountRepo *BankAccountPostgres) CreateAccount(bankIdentificationNum string, account entities.BankAccount) error {
	query := fmt.Sprintf("INSERT INTO %s (account_identif_num,bank_name,bank_identif_num) VALUES ($1,$2,$3)", AccountsTable)
	_, err := bankAccountRepo.db.Exec(query, account.AccountIdenitificationNum, account.BankFullName, account.BankIdentificationNum)
	return err
}

func (bankAccountRepo *BankAccountPostgres) FreezeBankAccount(accountIdenitificationNum string) error {
	query := fmt.Sprintf("UPDATE %s SET is_frozen=true WHERE account_identif_num=$1", AccountsTable)
	_, err := bankAccountRepo.db.Exec(query, accountIdenitificationNum)
	return err
}

func (bankAccountRepo *BankAccountPostgres) BlockBankAccount(accountIdenitificationNum string) error {
	query := fmt.Sprintf("UPDATE %s SET is_blocked=true WHERE account_identif_num=$1", AccountsTable)
	_, err := bankAccountRepo.db.Exec(query, accountIdenitificationNum)
	return err
}

func (bankAccountRepo *BankAccountPostgres) Transfer(sendAmount int, receiverAccountNum, senderAccountNum string) error {
	var moneyAmount int
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	checkAmountQuery := fmt.Sprintf("SELECT sendAmount FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(checkAmountQuery, senderAccountNum)
	err = row.Scan(&moneyAmount)
	if err != nil {
		tx.Rollback()
		return err
	}
	if moneyAmount < sendAmount {
		tx.Rollback()
		return errors.New("insufficient money amount")
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeSenderMoneyQuery, sendAmount, senderAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeReceiverMoneyQuery, sendAmount, receiverAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NewBankAccountPostgres(db *sql.DB) *BankAccountPostgres {
	return &BankAccountPostgres{db: db}
}
