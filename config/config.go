package config

import (
	"fmt"
	"os"

	"log"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
	"github.com/joho/godotenv"
)

var DB *sql.DB

func DB_Connnection() {
	//Loead Enviroment Varriable in go Program
	err := godotenv.Load()

	DB_USERNAME := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_NAME := os.Getenv("POSTGRES_DB_NAME")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection string
	dsn := "postgres://" + DB_USERNAME + ":" + DB_PASSWORD + "@localhost:5432/" + DB_NAME + "?sslmode=disable"

	// Open database connection
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL running in Docker!")

}
