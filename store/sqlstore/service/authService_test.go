package service_test

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/model"
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Create(t *testing.T){
	DatabaseURL := "host=localhost dbname = bmw-sales-test-db user = postgres password = asom sslmode=disable"
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.Service().AuthServiceInterface.Create(u))
	assert.NotNil(t,u)
}

func TestUserService_FindByPhone(t *testing.T) {
	DatabaseURL := "host=localhost dbname = bmw-sales-test-db user = postgres password = asom sslmode=disable"
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	phone := "998908080265"
	_,err := s.Service().AuthServiceInterface.FindByPhone(phone)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Phone = phone
	s.Service().AuthServiceInterface.Create(u)

	u,err = s.Service().AuthServiceInterface.FindByPhone(phone)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}