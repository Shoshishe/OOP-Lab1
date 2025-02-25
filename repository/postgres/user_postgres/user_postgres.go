package userPostgres

import "database/sql"

type UserPostgres struct {
	clientPostgres          ClientPostgres
	outerSpecialistPostgres OuterSpecialistPostgres
	managerPostgres         ManagerPostgres
	operatorPostgres        OperatorPostgres
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		clientPostgres:          *NewClientPostgres(db),
		outerSpecialistPostgres: *NewOuterSpecialistPostgres(db),
		managerPostgres:         *NewManagerPostgres(db),
		operatorPostgres:        *NewOperatorPostgres(db),
	}
}
