package infoPostgres

import (
	"database/sql"
	"fmt"
	"main/app_interfaces"
	"main/entities"
	"main/infrastructure/postgres"
)

type userInfoPostgres struct {
	app_interfaces.UserInfo
	db *sql.DB
}

func (userRepo *userInfoPostgres) GetRole(userId entities.UserRole) (entities.UserRole, error) {
	var role entities.UserRole
	query := fmt.Sprintf("SELECT role_id FROM %s WHERE id=$1 RETURNING role_id", postgres.UsersTable)
	res := userRepo.db.QueryRow(query, userId)
	err := res.Scan(&role)
	if err != nil {
		return 0, err
	}
	return role, nil
}

func NewUserInfoPostgres(db *sql.DB) *userInfoPostgres {
	return &userInfoPostgres{db: db}
}
