package postgres

import (
	"database/sql"
	"main/infrastructure"
)

type AuthPostgres struct {
	infrastructure.Authorization
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
