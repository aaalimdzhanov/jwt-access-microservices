package routes

import (
	"github.com/aaalimdzhanov/jwt-access-microservices/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes(storeRouter *mux.Router, h *controllers.Handler) *mux.Router {
	router := storeRouter.StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router,h)
	return router
}