package app

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	if os.Getenv("DB_USER") == "" {
		err := godotenv.Load("../.env")
		if err != nil {
			t.Fatal("Error loading .env file")
		}
	}
	db := NewDB()
	assert.NotNil(t, db, "database is null")
	err := db.Ping()
	assert.NoError(t, err, "Database connection should be successful")
	defer func() {
		err := db.Close()
		assert.NoError(t, err, "failed close database")
	}()
}
