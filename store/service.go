package store

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/model"
)

type AuthServiceInterface interface {
	Create(*model.User) error
	FindByPhone(phone string) (*model.User, error)
}

type Service struct {
	AuthServiceInterface
}

