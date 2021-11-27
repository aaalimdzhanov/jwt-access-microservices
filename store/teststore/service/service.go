package service

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/teststore/repository"
)

func NewService(repo *repository.Repository) store.Service {
	return store.Service{
		AuthServiceInterface: NewAuthService(repo.AuthRepositoryInterface),
	}
}
