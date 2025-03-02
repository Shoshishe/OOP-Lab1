package userPostgres

import "database/sql"

type UserPostgres struct {
	ClientPostgres
	OuterSpecialistPostgres
	ManagerPostgres
	OperatorPostgres
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		*NewClientPostgres(db),
		*NewOuterSpecialistPostgres(db),
		*NewManagerPostgres(db),
		*NewOperatorPostgres(db),
	}
}
