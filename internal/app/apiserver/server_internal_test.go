package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaalimdzhanov/jwt-access-microservices/model"
	"github.com/aaalimdzhanov/jwt-access-microservices/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T){
	s := NewServer(teststore.New())

	testCases :=[]struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"phone":"998909996019",
				"password":"password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"phone":"invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost,"/register",b)
			s.ServeHTTP(rec,req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionCreate(t *testing.T){

	u := model.TestUser(t)
	store := teststore.New()
	s := NewServer(store)
	store.Service().AuthServiceInterface.Create(u)

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		//{
		//	name: "valid",
		//	payload: map[string]string{
		//		"phone" : u.Phone,
		//		"password" : u.Password,
		//	},
		//	expectedCode: http.StatusOK,
		//},
		{
			name: "invalid payload",
			payload:"",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid phone",
			payload: map[string]string{
				"phone" : "invalid phone",
				"password" : u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost,"/login",b)
			s.ServeHTTP(rec,req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}