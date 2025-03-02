package persistanceMappers

import (
	"main/domain/entities"
	persistance "main/repositories/postgres/entities_models"
)

func ToUserEntity(usr persistance.UserPersistance, outsideInfo entities.UserOutside) (*entities.User, error) {
	return entities.NewUser(
		outsideInfo, usr.Password,
		usr.Email, entities.WithFullName(usr.FullName),
		entities.WithPasportNum(usr.PasportIdentifNum),
		entities.WithPasportSeries(usr.PasportSeries),
		entities.WithPhone(usr.MobilePhone),
		entities.WithUserRole(usr.RoleType),
		entities.WithUserId(usr.Id),
	)
}
