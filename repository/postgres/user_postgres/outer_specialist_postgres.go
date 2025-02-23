package userPostgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/repository/postgres"
	"main/service"
)

type OuterSpecialistPostgres struct {
	service.OuterSpecialistRepository
	db *sql.DB
}

func (repos *OuterSpecialistPostgres) SendInfoForPayment(req entities.PaymentRequest) error {
	query := fmt.Sprintf("INSERT INTO %s (client_id,account_num, amount, pasport_series, pasport_num) SELECT $1, $2, $3, pasport_series, pasport_num FROM %s", postgres.PaymentRequestsTable, postgres.UsersTable)
	_, err := repos.db.Exec(query, req.ClientId(), req.AccountNum(), req.Amount())
	if err != nil {
		return err
	}
	return nil
}

func (repos *OuterSpecialistPostgres) TransferRequest(transfer entities.Transfer) error {
	query := fmt.Sprintf("INSERT INTO %s (owner_id, sender_acc_num, receiver_acc_num) VALUES ($1,$2,$3)", postgres.TransfersTable)
	_, err := repos.db.Exec(query, transfer.TransferOwnerId(), transfer.SenderAccountNum(), transfer.ReceiverAccountNum())
	if err != nil {
		return err
	}
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
	row := repos.db.QueryRow(accountOwnerQuery, accountIdentifNum)
	err = row.Scan(&accountOwnerId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return false, err
		}
		return false, err
	}
	nonSharedCompaniesQuery := fmt.Sprintf("SELECT company_id FROM %s WHERE user_id=$1 AND company_id NOT IN (SELECT company_id FROM %s WHERE user_id = $2)", postgres.CompaniesWorkersTable, postgres.CompaniesWorkersTable)
	row = repos.db.QueryRow(nonSharedCompaniesQuery, accountOwnerId, specialistId)
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

func NewOuterSpecialistPostgres(db *sql.DB) *OuterSpecialistPostgres {
	return &OuterSpecialistPostgres{
		db: db,
	}
}
