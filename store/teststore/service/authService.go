package service

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/model"
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
)

type AuthService struct {
	repo store.AuthRepositoryInterface
}

func NewAuthService(repo store.AuthRepositoryInterface) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
func(s *AuthService) Create (u *model.User) error {
	return s.repo.Create(u)
}

func(s *AuthService) FindByPhone (phone string)(*model.User, error) {
	return s.repo.FindByPhone(phone)
}
