package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/service"
)

type BankAccountPostgres struct {
	service.AccountRepository
	db *sql.DB
}

func (bankAccountRepo *BankAccountPostgres) CreateAccount(account entities.BankAccount) error {
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

func (bankAccountRepo *BankAccountPostgres) PutMoney(amount entities.MoneyAmount, accountIdentificationNum string) error {
	query := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err := bankAccountRepo.db.Exec(query, amount, accountIdentificationNum)
	return err
}

func (bankAccountRepo *BankAccountPostgres) TakeMoney(amount entities.MoneyAmount, accountIdentificationNum string) error {
	query := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err := bankAccountRepo.db.Exec(query, amount, accountIdentificationNum)
	return err
}

func (bankAccountRepo *BankAccountPostgres) TransferMoney(transfer entities.Transfer) error {
	var moneyAmount int
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	checkAmountQuery := fmt.Sprintf("SELECT sendAmount FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(checkAmountQuery, transfer.SenderAccountNum)
	err = row.Scan(&moneyAmount)
	if err != nil {
		tx.Rollback()
		return err
	}
	if moneyAmount < int(transfer.SumOfTransfer) {
		tx.Rollback()
		return errors.New("insufficient money amount")
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeSenderMoneyQuery, transfer.SumOfTransfer, transfer.SenderAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeReceiverMoneyQuery, transfer.SumOfTransfer, transfer.ReceiverAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) CloseBankAccount(accountIdentifNum entities.AccountIdenitificationNum) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE account_identif_num=$1",AccountsTable)
	_, err := bankAccountRepo.db.Exec(query)
	if err != nil {
		return err
	}
	return err
}

func NewBankAccountPostgres(db *sql.DB) *BankAccountPostgres {
	return &BankAccountPostgres{db: db}
}
