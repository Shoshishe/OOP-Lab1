package app_interfaces

import "main/entities"

type TokenAuth interface {
	GenerateToken(fullName, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type UserInfo interface {
	GetRole(userId int) (entities.UserRole, error)
}