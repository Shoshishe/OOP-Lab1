package postgres

import (
	"database/sql"
	"fmt"
)

const (
	UsersTable        = "users"
	PendingUsersTable = "pending_users"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	DbName   string
	Password string
	SSLMode  string
}

func NewPostgresDb(conf DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host, conf.Port, conf.Username, conf.DbName, conf.Password, conf.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
