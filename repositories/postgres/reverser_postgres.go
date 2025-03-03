package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/domain/usecases"

	persistance "main/repositories/postgres/entities_models"
	persistanceMappers "main/repositories/postgres/mappers"

	"github.com/lib/pq"
)

type ReverserPostgres struct {
	usecases.ReverserInfo
	db *sql.DB
}

func NewReverserRepository(db *sql.DB) *ReverserPostgres {
	return &ReverserPostgres{
		db: db,
	}
}

func (repos *ReverserPostgres) GetAction(actionId int) (entities.Action, error) {
	var action persistance.ActionPersistance
	query := fmt.Sprintf("SELECT first_action_type, first_action_args FROM %s WHERE first_action_id = $1", ActionsTable)
	row := repos.db.QueryRow(query, actionId)
	err := row.Scan(&action.ActionName, pq.Array(&action.ActionArgs))
	if err != nil {
		if err == sql.ErrNoRows {
			query := fmt.Sprintf("SELECT second_action_type, second_action_args FROM %s WHERE second_action_id = $1", ActionsTable)
			row := repos.db.QueryRow(query, actionId)
			err := row.Scan(&action.ActionName, pq.Array(&action.ActionArgs))
			if err != nil {
				if err == sql.ErrNoRows {
					return entities.Action{}, errors.New("no action with such id")
				}
				return entities.Action{}, err
			}
		} else {
			return entities.Action{}, err
		}
	}
	return *persistanceMappers.ToActionEntity(action), nil
}
