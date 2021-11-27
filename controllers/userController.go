package controllers

import (
	"encoding/json"
	"errors"
	httpcontext "github.com/gorilla/context"
	"net/http"

	"github.com/aaalimdzhanov/jwt-access-microservices/common"
	"github.com/aaalimdzhanov/jwt-access-microservices/model"
)
var (
	errIncorectEmailOrPassword = errors.New("incorrect email or password")
)
func (h*Handler) HandleUsersCreate() http.HandlerFunc {
	type request struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil{
			common.Error(rw,r,http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Phone: req.Phone,
			Password: req.Password,
		}

		if err:= h.store.Service().AuthServiceInterface.Create(u); err != nil{
			common.Error(rw, r, http.StatusUnprocessableEntity,err)
			return
		}
		u.Sanitize()
		common.Respond(rw, r, http.StatusCreated, u)
	}
}

func(h*Handler) HandleSessionsCreate() http.HandlerFunc{
	type request struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil{
			common.Error(rw,r,http.StatusBadRequest, err)
			return
		}


		u, err := h.store.Service().AuthServiceInterface.FindByPhone(req.Phone)
		if err != nil || !u.ComparePassword(req.Password){
			common.Error(rw, r, http.StatusUnauthorized,errIncorectEmailOrPassword)
			return
		}
		token,_ := common.GenerateJWT(u)
		http.SetCookie(rw, &http.Cookie{
			Name:    "token",
			Value:   token,
		})
		common.Respond(rw, r, http.StatusOK, u)

	}
}

func (h*Handler) HandleWhoAmI() http.HandlerFunc {
	type request struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		if val, ok := httpcontext.GetOk(r, "user"); ok {

			common.Respond(rw, r, http.StatusCreated, val)
		}
	}
}
