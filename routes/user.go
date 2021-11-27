package routes

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/common"
	"github.com/aaalimdzhanov/jwt-access-microservices/controllers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router, h *controllers.Handler)  *mux.Router {
	router.HandleFunc("/register", h.HandleUsersCreate()).Methods("POST")
	router.HandleFunc("/login", h.HandleSessionsCreate()).Methods("POST")
	private := router.PathPrefix("/private").Subrouter()
	private.Use(common.AuthenticateUser)
	private.HandleFunc("/whoami", h.HandleWhoAmI()).Methods("POST")
	return router
}