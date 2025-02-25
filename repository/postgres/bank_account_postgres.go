package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/service/repository"
)

type BankAccountPostgres struct {
	repository.AccountRepository
	db *sql.DB
}

func (bankAccountRepo *BankAccountPostgres) CreateAccount(account entities.BankAccount, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (account_identif_num,bank_name,bank_identif_num) VALUES ($1,$2,$3)", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, account.AccountIdenitificationNum, account.BankFullName, account.BankIdentificationNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, account.AccountIdenitificationNum(), account.BankFullName(), account.BankIdentificationNum())
	err = InsertAction(tx, bankAccountRepo.db, "CreateAccount", args, usrId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (bankAccountRepo *BankAccountPostgres) ReverseAccountCreation(account entities.BankAccount, userId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE account_identif_num=$1", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, account.AccountIdenitificationNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, account.AccountIdenitificationNum(), account.BankFullName(), account.BankIdentificationNum())
	err = ReverseAction(tx, bankAccountRepo.db, args, userId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) FreezeBankAccount(accountIdenitificationNum string, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET is_frozen=true WHERE account_identif_num=$1", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, accountIdenitificationNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 1)
	args = append(args, accountIdenitificationNum)
	err = InsertAction(tx, bankAccountRepo.db, "FreezeBankAccount", args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) ReverseAccountFreeze(accountIdentifNum entities.AccountIdenitificationNum, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET is_frozen=false WHERE account_identif_num=$1", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, accountIdentifNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 1)
	args = append(args, accountIdentifNum)
	err = ReverseAction(tx, bankAccountRepo.db, args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) BlockBankAccount(accountIdenitificationNum string, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET is_blocked=true WHERE account_identif_num=$1", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, accountIdenitificationNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 1)
	args = append(args, accountIdenitificationNum)
	err = InsertAction(tx, bankAccountRepo.db, repository.BlockAccountAction, args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) ReverseAccountBlock(accountIdentifNum entities.AccountIdenitificationNum, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET is_blocked=false WHERE account_identif_num=$1", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, accountIdentifNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 1)
	args = append(args, accountIdentifNum)
	err = ReverseAction(tx, bankAccountRepo.db, args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
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

func (bankAccountRepo *BankAccountPostgres) TransferMoney(transfer entities.Transfer, usrId int) error {
	var moneyAmount int
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	checkAmountQuery := fmt.Sprintf("SELECT amount FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(checkAmountQuery, transfer.SenderAccountNum())
	err = row.Scan(&moneyAmount)
	if err != nil {
		tx.Rollback()
		return err
	}
	if moneyAmount < int(transfer.SumOfTransfer()) {
		tx.Rollback()
		return errors.New("insufficient money amount")
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeSenderMoneyQuery, transfer.SumOfTransfer(), transfer.SenderAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeReceiverMoneyQuery, transfer.SumOfTransfer(), transfer.ReceiverAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), fmt.Sprint(transfer.SumOfTransfer()))
	err = InsertAction(tx, bankAccountRepo.db, "TransferMoney", args, usrId)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) ReverseMoneyTransfer(transfer entities.Transfer, usrId int) error {
	var moneyAmount int
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	checkAmountQuery := fmt.Sprintf("SELECT sendAmount FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(checkAmountQuery, transfer.SenderAccountNum())
	err = row.Scan(&moneyAmount)
	if err != nil {
		tx.Rollback()
		return err
	}
	if moneyAmount < int(transfer.SumOfTransfer()) {
		tx.Rollback()
		return errors.New("insufficient money amount")
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeSenderMoneyQuery, transfer.SumOfTransfer(), transfer.SenderAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(changeReceiverMoneyQuery, transfer.SumOfTransfer(), transfer.ReceiverAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}

	args := make([]string, 0, 3)
	args = append(args, transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), fmt.Sprint(transfer.SumOfTransfer()))
	err = ReverseAction(tx, bankAccountRepo.db, args, usrId)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) CloseBankAccount(accountIdentifNum entities.AccountIdenitificationNum, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE account_identif_num=$1", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	// actionInsertQuery := fmt.Sprintf("INSERT INTO %s (user_id, first_action_type, first_action_args) VALUES ($1,$2,$3) ON CONFLICT (user_id) DO UPDATE SET"+
	// 	"second_action_type=first_action_type, second_action_args=first_action_args,"+
	// 	"first_action_type=EXCLUDED.first_action_type, first_action_args=EXCLUDED.first_action_args", ActionsTable)
	// args := make([]string, 0, 1)
	// args = append(args, accountIdentifNum)
	// _, err = bankAccountRepo.db.Exec(actionInsertQuery, usrId, "CloseBankAccount", pq.Array(args))
	
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	tx.Commit()
	return err
}

func NewBankAccountPostgres(db *sql.DB) *BankAccountPostgres {
	return &BankAccountPostgres{db: db}
}
