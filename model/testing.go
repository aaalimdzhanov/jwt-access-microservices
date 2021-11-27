package model

import "testing"

func TestUser(t *testing.T) *User{
	return &User{
		Phone: "998909996019",
		Password: "password",
	}
}