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

func (repos *ManagerPostgres) CancelOuterWorkerOperation(operationId int) error{
	//DO THE WHOLE SHIT
	return nil
}

func (repos *ManagerPostgres) ApproveCredit(requestId int) error {
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
		db: db,
	}
}