package persistanceMappers

import (
	"main/domain/entities"
	persistance "main/repositories/postgres/entities_models"
)

func ToActionEntity(actionPersistance persistance.ActionPersistance) *entities.Action {
	return entities.NewAction(actionPersistance.ActionId, actionPersistance.ActionName, actionPersistance.ActionArgs)
}
