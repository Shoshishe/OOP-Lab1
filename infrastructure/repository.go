package infrastructure

import "main/entities"

type Authorization interface {
	AddUser(user entities.User) (int, error)
	GetUser(fullName, password string) (*entities.User, error)
}

type Repository struct {
	Authorization
}


