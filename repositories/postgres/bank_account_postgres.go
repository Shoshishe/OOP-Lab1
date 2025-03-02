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

type BankAccountPostgres struct {
	repository.AccountRepository
	db *sql.DB
}

func (bankAccountRepo *BankAccountPostgres) GetAccounts(usrId int) ([]entities.BankAccount, error) {
	var accountList []entities.BankAccount
	query := fmt.Sprintf("SELECT bank_name,account_identif_num,amount,is_frozen,bank_identif_num FROM %s WHERE user_id=$1", AccountsTable)
	rows, err := bankAccountRepo.db.Query(query, usrId)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.New("no accounts to show :(")
	}
	account := persistance.BankAccount{}
	err = rows.Scan(&account.BankFullName, &account.AccountIdenitificationNum, &account.Amount,
		&account.IsFrozen, &account.BankIdentificationNum)
	if err != nil {
		return nil, err
	}
	tempAccount, err := persistanceMappers.ToAccountEntity(&account, entities.NewResponseValidatePolicy())
	if err != nil {
		return nil, err
	}
	accountList = append(accountList, *tempAccount)
	i := 1
	for rows.Next() {
		err = rows.Scan(&account.BankFullName, &account.AccountIdenitificationNum, &account.Amount,
			&account.IsFrozen, &account.BankFullName)
		if err != nil {
			return nil, err
		}
		account, err := persistanceMappers.ToAccountEntity(&account, entities.NewResponseValidatePolicy())
		accountList = append(accountList, *account)
		if err != nil {
			return nil, err
		}
		i++
	}
	return accountList, nil
}
func (bankAccountRepo *BankAccountPostgres) CreateAccountAsPerson(account entities.BankAccount, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (account_identif_num,bank_name,bank_identif_num, user_id) VALUES ($1,$2,$3, $4)", AccountsTable)
	_, err = tx.Exec(query, account.AccountIdenitificationNum(), account.BankFullName(), account.BankIdentificationNum(), usrId)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 4)
	args = append(args, account.AccountIdenitificationNum(), account.BankFullName(), account.BankIdentificationNum(), fmt.Sprint(usrId))
	err = InsertAction(tx, "CreateAccountAsUser", args, usrId)
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

func (bankAccountRepo *BankAccountPostgres) CreateAccountAsCompany(account entities.BankAccount, companyId int, issuerId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (account_identif_num,bank_name,bank_identif_num, company_id, user_id) VALUES ($1,$2,$3,$4,$5)", AccountsTable)
	_, err = tx.Exec(query, account.AccountIdenitificationNum(), account.BankFullName(), account.BankIdentificationNum(), companyId, issuerId)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 5)
	args = append(args, account.AccountIdenitificationNum(), account.BankFullName(), account.BankIdentificationNum(), fmt.Sprint(companyId), fmt.Sprint(issuerId))
	err = InsertAction(tx, "CreateAccountAsCompany", args, issuerId)
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
	err = InsertAction(tx, "FreezeBankAccount", args, usrId)
	if err != nil {
		tx.Rollback()
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
	err = InsertAction(tx, repository.BlockAccountAction, args, usrId)
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
	var accountAmount int
	checkAmount := fmt.Sprintf("SELECT amount from %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(checkAmount, accountIdentificationNum)
	err := row.Scan(&accountAmount)
	if err != nil {
		return err
	}
	if accountAmount < int(amount) {
		return errors.New("not enough money on the account")
	}
	query := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = bankAccountRepo.db.Exec(query, amount, accountIdentificationNum)
	return err
}

func (bankAccountRepo *BankAccountPostgres) TransferMoney(transfer entities.Transfer, usrId int) error {
	var moneyAmount int
	var isFrozen bool
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	checkAmountAndFrozenQuery := fmt.Sprintf("SELECT amount, is_frozen FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := tx.QueryRow(checkAmountAndFrozenQuery, transfer.SenderAccountNum())
	err = row.Scan(&moneyAmount, &isFrozen)
	if err != nil {
		tx.Rollback()
		return err
	}
	if moneyAmount < int(transfer.SumOfTransfer()) {
		tx.Rollback()
		return errors.New("insufficient money amount")
	}
	if isFrozen {
		tx.Rollback()
		return errors.New("account is frozen")
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = tx.Exec(changeSenderMoneyQuery, transfer.SumOfTransfer(), transfer.SenderAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = tx.Exec(changeReceiverMoneyQuery, transfer.SumOfTransfer(), transfer.ReceiverAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), fmt.Sprint(transfer.SumOfTransfer()))
	err = InsertAction(tx, "TransferMoney", args, usrId)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (bankAccountRepo *BankAccountPostgres) ReverseMoneyTransfer(transfer entities.Transfer, usrId int) error {
	tx, err := bankAccountRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = tx.Exec(changeSenderMoneyQuery, transfer.SumOfTransfer(), transfer.SenderAccountNum())
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", AccountsTable)
	_, err = tx.Exec(changeReceiverMoneyQuery, transfer.SumOfTransfer(), transfer.ReceiverAccountNum())
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
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (bankAccountRepo *BankAccountPostgres) DoesBankExist(bankIdentifNum entities.BankIdentificationNum, bankFullName entities.BankName) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE bank_ident_num=$1 AND name=$2", BanksTable)
	_, err := bankAccountRepo.db.Query(query, bankIdentifNum, bankFullName)
	if err != nil {
		return false, err
	}
	if err := bankAccountRepo.db.QueryRow(query, bankIdentifNum, bankFullName).Scan(&exists); err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return false, err
	}
}

func (bankAccountRepo *BankAccountPostgres) DoesAccountBelongToUser(accountNum string, userId int) (bool, error) {
	var doesBelong bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE account_identif_num=$1 AND user_id=$2", AccountsTable)
	row := bankAccountRepo.db.QueryRow(query, accountNum, userId)
	err := row.Scan(&doesBelong)
	if err != nil {
		return false, err
	}
	return doesBelong, nil
}

func (bankAccountRepo *BankAccountPostgres) DoesAccountExist(accountNum string) (bool, error) {
	var doesBelong bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(query, accountNum)
	err := row.Scan(&doesBelong)
	if err != nil {
		return false, err
	}
	return doesBelong, nil
}

func (bankAccountRepo *BankAccountPostgres) AccountMoneyAmount(accountNum entities.AccountIdenitificationNum) (entities.MoneyAmount, error) {
	var amount int
	query := fmt.Sprintf("SELECT amount FROM %s WHERE account_identif_num=$1", AccountsTable)
	row := bankAccountRepo.db.QueryRow(query, accountNum)
	err := row.Scan(&amount)
	if err != nil {
		return 0, err
	}
	return entities.MoneyAmount(amount), nil
}
func NewBankAccountPostgres(db *sql.DB) *BankAccountPostgres {
	return &BankAccountPostgres{db: db}
}
