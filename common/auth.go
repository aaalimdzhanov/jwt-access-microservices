package common

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aaalimdzhanov/jwt-access-microservices/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
)

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)
const (
	// JWT Record and claim
	JWT_COL_NAME         = "jwt"
	JWT_USER_ATTR        = "user"
	JWT_PASS_ATTR        = "pass"
	JWT_ENDPOINTS_ATTR   = "endpoints"
	JWT_COLLECTIONS_ATTR = "collections"
	JWT_USER_ADMIN       = "admin"
	// JWT claim
	JWT_EXPIRY = "exp"
)
type AppClaims struct {
	jwt.StandardClaims
	Phone string `json:"phone"`
}
// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)
// Read the key files before starting http handlers
func InitKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		fmt.Println(privKeyPath)
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}
// Generate JWT token
func GenerateJWT(user *model.User) (string, error) {

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = jwt.MapClaims{
		JWT_USER_ATTR:        user.Phone,
		JWT_EXPIRY:           time.Now().Add(time.Hour * 72).Unix(),
	}
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

// Extract JWT from Authorization header or "access_token" attribute.
type TokenExtractor struct {
}

func (t TokenExtractor) ExtractToken(req *http.Request) (string, error) {
	token := req.Header.Get("Authorization")
	if token == "" {
		token = req.FormValue("access_token")
	}
	if token == "" {
		return "", request.ErrNoTokenInRequest
	}
	// For the sake of simplicity, extra spaces and type name Bearer are stripped.
	return strings.TrimSpace(strings.TrimPrefix(token, "Bearer")), nil
}


func  AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from request
		token, err := request.ParseFromRequest(r, TokenExtractor{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return verifyKey, nil
		})
		if err != nil {
			Error(w, r, http.StatusInternalServerError, err)
			return
		}
		if !token.Valid {
			Error(w, r, http.StatusInternalServerError, err)
			return
		}
		tokenClaims := token.Claims.(jwt.MapClaims)
		context.Set(r, "user", tokenClaims[JWT_USER_ATTR])
		next.ServeHTTP(w, r)
	})
}




