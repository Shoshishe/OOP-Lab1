package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/entities"
	"main/infrastructure"
)

type AuthPostgres struct {
	infrastructure.Authorization
	db *sql.DB
}

func (authRepo *AuthPostgres) AddUser(user entities.User) (int, error) {
	var id int
	tx, err := authRepo.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s (full_name, pasport_series, phone_number, email, password, role_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING ID", UsersTable)
	row := authRepo.db.QueryRow(query,user.FullName,user.PasportSeries,user.MobilePhone,user.Email,user.Password,user.RoleType)
	if err := row.Scan(&id); err != nil {
		rollbackErr := tx.Rollback()
		err = errors.Join(rollbackErr, err)
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
