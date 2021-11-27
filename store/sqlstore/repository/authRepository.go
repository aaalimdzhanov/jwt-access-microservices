package repository

import (
	"database/sql"
	"github.com/aaalimdzhanov/jwt-access-microservices/model"
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	Db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository{
	return &AuthRepository{
		Db: db,
	}
}

func (r *AuthRepository) Create(u *model.User) error  {
	if err := u.Validate(); err != nil{
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.Db.QueryRow("INSERT INTO USERS(phone, encrypted_password) VALUES($1, $2) RETURNING Id",
		u.Phone,
		u.EncryptedPassword,
	).Scan(&u.Id)
}

func(r *AuthRepository) FindByPhone (phone string)(*model.User, error) {
	u := &model.User{}
	if err := r.Db.QueryRow("SELECT id,phone,encrypted_password from users where phone = $1", phone).Scan(
		&u.Id,
		&u.Phone,
		&u.EncryptedPassword,
	);
		err == sql.ErrNoRows{
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
