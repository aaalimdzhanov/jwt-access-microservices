package store

import "github.com/aaalimdzhanov/jwt-access-microservices/model"

type AuthRepositoryInterface interface {
	Create(*model.User) error
	FindByPhone(phone string) (*model.User, error)
}