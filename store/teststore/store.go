package teststore

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/teststore/repository"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/teststore/service"
)

type Store struct {

	service store.Service
}

func New() *Store {
	return &Store{

	}
}

func (s*Store) Service() store.Service  {
	repo := repository.NewRepository()
	s.service = service.NewService(repo)
	return  s.service
}
