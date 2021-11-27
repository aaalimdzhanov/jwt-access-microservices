package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/aaalimdzhanov/jwt-access-microservices/controllers"
	"github.com/aaalimdzhanov/jwt-access-microservices/routes"
	"github.com/aaalimdzhanov/jwt-access-microservices/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)
var (
	errIncorectEmailOrPassword = errors.New("incorrect email or password")
)
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

func NewServer(store store.Store) *server{
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}
	
	s.configureRouter()
	return s
}

func(s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	 s.router.ServeHTTP(w,r)
}

func(s *server) configureRouter(){
	handler := controllers.NewController(s.store)
	routes.InitRoutes(s.router, handler)
}
func(s *server) error(rw http.ResponseWriter, r *http.Request, code int, err error){
	s.respond(rw, r, code, map[string]string{"error":err.Error()})
}

func(s *server) respond(rw http.ResponseWriter, r *http.Request, code int, data interface{}){
	rw.WriteHeader(code)
	if data != nil{
		json.NewEncoder(rw).Encode(data)
	}
}