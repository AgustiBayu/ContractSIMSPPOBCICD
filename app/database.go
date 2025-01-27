package app

import (
	"ContractSIMSPPOBCICD/helper"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbSslmode := os.Getenv("DB_SSLMODE")

	fmt.Println("DB_USER:", dbUser)
	fmt.Println("DB_PASS:", dbPass)
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_NAME:", dbName)
	fmt.Println("DB_SSLMODE:", dbSslmode)

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" || dbSslmode == "" {
		panic("Missing required database environment variables")
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbName, dbSslmode)
	db, err := sql.Open("postgres", dsn)
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
