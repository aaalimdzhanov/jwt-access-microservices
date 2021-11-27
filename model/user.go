package model

import (
	"regexp"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int `json:"id"`
	Phone             string `json:"phone"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil{
		return "", err
	}
	return string(b), nil
}

func (u *User) Validate() error{
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Phone, validation.Required, validation.Match(regexp.MustCompile("(998)[0-9]{9}"))),
		validation.Field(&u.Password,validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6,100)),

	)
}

func (u *User) Sanitize(){
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool{
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}