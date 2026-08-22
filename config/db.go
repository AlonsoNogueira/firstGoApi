package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetUpDataBase() *sql.DB {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DB_CONNECTION")

	fmt.Printf("%q\n", connectionString)

	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connect to database")

	return dbConnection
}
