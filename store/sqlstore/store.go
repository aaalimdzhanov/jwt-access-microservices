package sqlstore

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/sqlstore/repository"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/sqlstore/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	Db    *sqlx.DB
	service store.Service
}

func New(db *sqlx.DB) *Store {
	return &Store{
		Db: db,
	}
}

func (s*Store) Service() store.Service  {

	repo := repository.NewRepository(s.Db)
	s.service = service.NewService(repo)
	return  s.service
}

