package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	persistance "main/repositories/postgres/entities_models"
	persistanceMappers "main/repositories/postgres/mappers"
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
	query := fmt.Sprintf("INSERT INTO %s (full_name, pasport_series, phone_number, email, password, role_id, pasport_num) VALUES ($1,$2,$3,$4,$5,$6, $7) RETURNING ID", UsersTable)
	row := tx.QueryRow(query, user.FullName(), user.PasportSeries(), user.MobilePhone(), user.Email(), user.Password(), user.RoleType(), user.PasportNum())
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

func (authRepo *AuthPostgres) GetUser(email, password string) (*entities.User, error) {
	var persistanceUsr persistance.UserPersistance

	query := fmt.Sprintf("SELECT id, full_name, pasport_series, phone_number, email, password, role_id, pasport_num FROM %s WHERE email=$1 and password=$2", UsersTable)
	row := authRepo.db.QueryRow(query, email, password)
	if err := row.Scan(&persistanceUsr.Id, &persistanceUsr.FullName, &persistanceUsr.PasportSeries,
		&persistanceUsr.MobilePhone, &persistanceUsr.Email, &persistanceUsr.Password,
		&persistanceUsr.RoleType, &persistanceUsr.PasportIdentifNum); err != nil {
		return &entities.User{}, err
	}

	return persistanceMappers.ToUserEntity(persistanceUsr, authRepo)
}

func (authRepo *AuthPostgres) GetUserRole(userId int) (entities.UserRole, error) {
	var usrRole entities.UserRole
	query := fmt.Sprintf("SELECT role_id FROM %s WHERE id=$1", UsersTable)
	row := authRepo.db.QueryRow(query, userId)
	if err := row.Scan(&usrRole); err != nil {
		return 0, err
	}
	return usrRole, nil
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
