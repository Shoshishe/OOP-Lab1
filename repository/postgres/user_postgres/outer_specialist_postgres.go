package userPostgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/repository/postgres"
	"main/service/repository"
)

type OuterSpecialistPostgres struct {
	repository.OuterSpecialistRepository
	db *sql.DB
}

func (repos *OuterSpecialistPostgres) SendInfoForPayment(req entities.PaymentRequest, usrId int) error {
	tx, err := repos.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (client_id,account_num, amount, pasport_series, pasport_num) SELECT $1, $2, $3, pasport_series, pasport_num FROM %s WHERE id=$1", postgres.PaymentRequestsTable, postgres.UsersTable)
	_, err = tx.Exec(query, req.ClientId(), req.AccountNum(), req.Amount())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprint(req.ClientId()), req.AccountNum(), fmt.Sprint(req.Amount()))
	err = postgres.InsertAction(tx,repository.SendInfoForPaymentAction, args, usrId)
	if err != nil {
		return err
	}
    tx.Commit()
	return nil
}

func (repos *OuterSpecialistPostgres) TransferRequest(transfer entities.Transfer, usrId int) error {
	tx, err := repos.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (owner_id, sender_acc_num, receiver_acc_num, amount) VALUES ($1,$2,$3, $4)", postgres.TransfersTable)
	_, err = tx.Exec(query, transfer.TransferOwnerId(), transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), transfer.SumOfTransfer())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 4)
	args = append(args, fmt.Sprint(transfer.TransferOwnerId()),transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), fmt.Sprint(transfer.SumOfTransfer()))
	err = postgres.InsertAction(tx, repository.TransferRequestAction, args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (repos *OuterSpecialistPostgres) ReverseTransferRequest(transfer entities.Transfer, usrId int) error {
	tx, err := repos.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE owner_id=$1 AND sender_acc_num=$2 AND receiver_acc_num=$3 AND amount=$4", postgres.TransfersTable)
	_, err = tx.Exec(query, transfer.TransferOwnerId(), transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), transfer.SumOfTransfer())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 4)
	args = append(args, fmt.Sprint(transfer.TransferOwnerId()),transfer.SenderAccountNum(), transfer.ReceiverAccountNum(), fmt.Sprint(transfer.SumOfTransfer()))
	err = postgres.ReverseAction(tx, repos.db,  args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (repos *OuterSpecialistPostgres) DoesAccountBelongToOuterCompany(accountNum entities.AccountIdenitificationNum, specialistId int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s acd WHERE acd.account_identif_num=$1 INNER JOIN %s cd ON cd.id <> acd.company_id)", postgres.AccountsTable, postgres.CompaniesTable)
	row := repos.db.QueryRow(query, specialistId, accountNum)
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repos *OuterSpecialistPostgres) DoesAccountBelongToNonOuterUser(accountIdentifNum entities.AccountIdenitificationNum, specialistId int) (bool, error) {
	var accountOwnerId int
	tx, err := repos.db.Begin()
	if err != nil {
		return false, err
	}
	accountOwnerQuery := fmt.Sprintf("SELECT 1 FROM %s acd WHERE acd.account_identif_num=$1 INNER JOIN %s ud ON ud.id=acd.user_id RETURNING ud.id", postgres.AccountsTable, postgres.UsersTable)
	row := tx.QueryRow(accountOwnerQuery, accountIdentifNum)
	err = row.Scan(&accountOwnerId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return false, err
		}
		return false, err
	}
	nonSharedCompaniesQuery := fmt.Sprintf("SELECT company_id FROM %s WHERE user_id=$1 AND company_id NOT IN (SELECT company_id FROM %s WHERE user_id = $2)", postgres.CompaniesWorkersTable, postgres.CompaniesWorkersTable)
	row = tx.QueryRow(nonSharedCompaniesQuery, accountOwnerId, specialistId)
	if row.Err() == sql.ErrNoRows {
		return false, errors.New("no distinct companies between users")
	}
	if row.Err() != nil {
		return false, err
	}
	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repos *OuterSpecialistPostgres) DoesAccountBelongToUser(accountNum entities.AccountIdenitificationNum, userId int) (bool, error) {
	var belongs bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE acount_identif_num=$1 AND user_id=$2)", postgres.AccountsTable)
	row := repos.db.QueryRow(query, accountNum, userId)
	err := row.Scan(&belongs)
	if err != nil {
		return false, err
	}
	return belongs, nil
}

func (repos *OuterSpecialistPostgres) AccountMoneyAmount(accountNum entities.AccountIdenitificationNum) (entities.MoneyAmount, error) {
	var moneyAmount int
	query := fmt.Sprintf("SELECT amount FROM %s WHERE acount_identif_num=$1", postgres.AccountsTable)
	row := repos.db.QueryRow(query, accountNum)
	err := row.Scan(moneyAmount)
	if err != nil {
		return 0, err
	}
	return entities.MoneyAmount(moneyAmount), nil
}

func NewOuterSpecialistPostgres(db *sql.DB) *OuterSpecialistPostgres {
	return &OuterSpecialistPostgres{
		db: db,
	}
}
