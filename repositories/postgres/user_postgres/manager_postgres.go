package userPostgres

import (
	"database/sql"
	"fmt"
	"main/repositories/postgres"
	"main/service/repository"
)

type ManagerPostgres struct {
	operator OperatorPostgres
	repository.ManagerRepository
	db *sql.DB
}

func (repos *ManagerPostgres) ApproveCredit(requestId int, usrId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_approved=true WHERE request_id=$1", postgres.LoansTable)
	_, err := repos.db.Exec(query, requestId)
	if err != nil {
		return err
	}
	return nil
}

func NewManagerPostgres(db *sql.DB) *ManagerPostgres {
	return &ManagerPostgres{
		operator: *NewOperatorPostgres(db),
		db:       db,
	}
}
