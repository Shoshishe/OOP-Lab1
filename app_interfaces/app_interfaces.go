package app_interfaces

type TokenAuth interface {
	GenerateToken(fullName, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type UserInfo interface {
	GetRole(userId int) (int, error)
}