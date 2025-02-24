package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/service/repository"
)

type AuthPostgres struct {
	repository.AuthorizationRepository
	db *sql.DB
}

func (authRepo *AuthPostgres) AddUser(user entities.User) error {
	var id int
	tx, err := authRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (full_name, pasport_series, phone_number, email, password, role_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING ID", UsersTable)
	row := authRepo.db.QueryRow(query, user.FullName(), user.PasportSeries(), user.MobilePhone(), user.Email(), user.Password(), user.RoleType())
	if err := row.Scan(&id); err != nil {
		rollbackErr := tx.Rollback()
		err = errors.Join(rollbackErr, err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (authRepo *AuthPostgres) GetUser(fullName, password string) (*entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE full_name=$1 and password=$2", UsersTable)
	row := authRepo.db.QueryRow(query, fullName, password)
	if err := row.Scan(&user); err != nil {
		return &entities.User{}, err
	}
	return &user, nil
}

func (authRepo *AuthPostgres) GetRole(userId int) (entities.UserRole, error) {
	return entities.RolePendingUser, nil
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
