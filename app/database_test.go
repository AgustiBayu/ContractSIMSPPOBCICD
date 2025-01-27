package app

import (
	"ContractSIMSPPOBCICD/helper"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	err := godotenv.Load("../.env")
	helper.PanicIfError(err)
	db := NewDB()
	assert.NotNil(t, db, "database is null")
	err = db.Ping()
	assert.NoError(t, err, "Database connection should be successful")
	defer func() {
		err := db.Close()
		assert.NoError(t, err, "failed close database")
	}()
}
