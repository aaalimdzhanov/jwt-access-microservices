package repository

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/model"
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
)

type AuthRepository struct {
	users map[string]*model.User
}

func NewAuthRepository() *AuthRepository{
	return &AuthRepository{
		users: make(map[string]*model.User),
	}
}

func (r *AuthRepository) Create(u *model.User) error  {
	if err := u.Validate(); err != nil{
		return  err
	}

	if err := u.BeforeCreate(); err != nil{
		return  err
	}

	r.users[u.Phone] = u
	u.Id = len(r.users)

	return nil
}

func(r *AuthRepository) FindByPhone (phone string)(*model.User, error) {
	u, ok := r.users[phone]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
