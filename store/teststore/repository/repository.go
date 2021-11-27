package repository

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/store"

)

type Repository struct {
	store.AuthRepositoryInterface
}

func NewRepository() *Repository  {
	return &Repository{
		AuthRepositoryInterface:NewAuthRepository(),
	}
}