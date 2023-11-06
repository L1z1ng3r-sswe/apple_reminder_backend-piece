package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var DB *sql.DB

func InitDB() error {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
		return err
	}

	log.Println("Connected to the database")
	return nil
}

func CloseDB() {
	DB.Close()
}