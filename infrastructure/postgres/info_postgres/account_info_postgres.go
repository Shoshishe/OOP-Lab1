package infoPostgres

import (
	"database/sql"
	"fmt"
	"main/app_interfaces"
	"main/infrastructure/postgres"
)

type accountInfoPostgres struct {
	app_interfaces.AccountInfo
	db *sql.DB
}

func (accountInfoRepos *accountInfoPostgres) CheckBelonging(userId int, accountIdentifNum string) (bool, error) {
	query := fmt.Sprintf("SELECT FROM %s WHERE user_id=$1, account_identif_num=$2", postgres.AccountsTable)
	row := accountInfoRepos.db.QueryRow(query, userId, accountIdentifNum)
	if row.Err() == sql.ErrNoRows {
		return false, nil
	} else if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

func NewAccountInfoPostgres(db *sql.DB) *accountInfoPostgres {
	return &accountInfoPostgres{db: db}
}
