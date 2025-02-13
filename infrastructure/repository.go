package infrastructure

import "main/entities"

type Authorization interface {
	AddUser(user entities.User) (int, error)
	AddForApproval(user entities.User) (int, error)
	GetUser(fullName, password string) (*entities.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(authRepos Authorization) *Repository {
	return &Repository{
		Authorization: authRepos,
	}
}
