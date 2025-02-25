package userPostgres

import (
	"database/sql"
	"fmt"
	"main/repository/postgres"
	"main/service/repository"
)

type ManagerPostgres struct {
	operator OperatorPostgres
	repository.ManagerRepository
	db *sql.DB
}

func (repos *ManagerPostgres) CancelOuterWorkerOperation(operationId int, userId int) error {
	//DO THE WHOLE SHIT
	return nil
}

func (repos *ManagerPostgres) ApproveCredit(requestId int, usrId int) error {
	tx, err := repos.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET is_approved=true WHERE request_id=$1", postgres.LoansTable)
	_, err = repos.db.Exec(query, requestId)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprint(requestId))
	err = postgres.InsertAction(tx, repos.db, "ApproveCredit", args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func NewManagerPostgres(db *sql.DB) *ManagerPostgres {
	return &ManagerPostgres{
		operator: *NewOperatorPostgres(db),
		db:       db,
	}
}
