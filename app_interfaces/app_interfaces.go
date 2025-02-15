package app_interfaces

import "main/entities"

type TokenAuth interface {
	GenerateToken(fullName, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Info interface {
	UserInfo
	AccountInfo
}

type UserInfo interface {
	GetUserRole(userId int) (entities.UserRole, error)
}
type AccountInfo interface {
	CheckBelonging(userId int, accountIdentifNum string) (bool, error)
}