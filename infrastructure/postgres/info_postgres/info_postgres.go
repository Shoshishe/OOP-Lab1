package infoPostgres

import (
	"database/sql"
	"main/app_interfaces"
)

type InfoPostgres struct {
	app_interfaces.Info
	userInfoRepos    *userInfoPostgres
	accountInfoRepos *accountInfoPostgres
	db               *sql.DB
}

func NewInfoPostgres(db *sql.DB) *InfoPostgres {
	return &InfoPostgres{
		db:            db,
		userInfoRepos: NewUserInfoPostgres(db),
		accountInfoRepos: NewAccountInfoPostgres(db),
	}
}
