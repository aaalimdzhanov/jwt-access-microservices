package repository

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	store.AuthRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository  {
	return &Repository{
		AuthRepositoryInterface:NewAuthRepository(db),
	}
}