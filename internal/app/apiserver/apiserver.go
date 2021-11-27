package apiserver

import (
	"net/http"

	"github.com/aaalimdzhanov/jwt-access-microservices/store/sqlstore"
	"github.com/jmoiron/sqlx"
)

func Start(config *Config) error {

	db, err := newDb(config.DatabaseURL)
	if err != nil{
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := NewServer(store)



	return http.ListenAndServe(config.BindAdr, srv)
}

func newDb(databaseUrl string) (*sqlx.DB, error){
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil{
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}