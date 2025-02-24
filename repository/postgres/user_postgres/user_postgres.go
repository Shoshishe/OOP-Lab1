package userPostgres

import "database/sql"

type UserPostgres struct {
	clientPostgres          ClientPostgres
	outerSpecialistPostgres OuterSpecialistPostgres
	managerPostgres ManagerPostgres
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		clientPostgres: *NewClientPostgres(db),
		outerSpecialistPostgres: *NewOuterSpecialistPostgres(db),
		managerPostgres: *NewManagerPostgres(db),
	}
}
