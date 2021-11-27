package controllers

import "github.com/aaalimdzhanov/jwt-access-microservices/store"

type Handler struct {
	store store.Store
}
func NewController(store store.Store) *Handler{
	return &Handler{
		store: store,
	}
}