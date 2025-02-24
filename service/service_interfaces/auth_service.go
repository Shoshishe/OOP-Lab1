package serviceInterfaces

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	"main/service/entities_models/response"
)

type RoleAccess interface {
	GetUserRole(userId int) (entities.UserRole, error)
}
type TokenAuth interface {
	GenerateToken(fullName, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	RoleAccess
}
type Authorization interface {
	AddUser(user request.ClientSignUpModel) error
	// AddAdmin(admin request.AdminSignUpModel) error
	// AddManager(manager request.ManagerSignUpModel) error 
	// AddOuterSpecialist(manager request.ManagerSignUpModel) error
	GetUser(username, password string) (*response.UserAuthModel, error)
}