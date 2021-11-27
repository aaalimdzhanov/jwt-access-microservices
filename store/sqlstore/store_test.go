package sqlstore_test

import (
	"os"
	"testing"
)

var (
	DatabaseURL string
)
func TestMain(m *testing.M){
	DatabaseURL := os.Getenv("DATABASE_URL")
	if DatabaseURL == ""{
		DatabaseURL = "host=localhost dbname = bmw-test-store user = postgres password = asom sslmode=disable"
	}

	os.Exit(m.Run())
}