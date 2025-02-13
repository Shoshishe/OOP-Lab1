package postgres

import (
	"database/sql"
	"main/entities"
	"main/infrastructure"
)

type AuthPostgres struct {
	infrastructure.Authorization
	db *sql.DB
}

func (*AuthPostgres) AddForApproval(user entities.User) (int,error) {
	//query := fmt.Sprintf("INSERT INTO %s (full_name, pasport_series, phone_number, email, password, role_id) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING ID", PendingUsersTable)
	return 0, nil 
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
