package app

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	// Coba load file .env jika ada di direktori yang lebih tinggi
	if _, err := os.Stat("../.env"); err == nil {
		// Jika file ada, load file .env
		err := godotenv.Load("../.env")
		if err != nil {
			t.Fatal("Error loading .env file")
		}
	} else {
		// Jika file .env tidak ditemukan, log pesan tetapi lanjutkan untuk menggunakan environment variables dari sistem
		t.Log("Warning: .env file not found, using system environment variables")
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
