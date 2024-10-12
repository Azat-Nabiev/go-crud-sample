package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func main() {
	database := initDB()
	defer database.Close()
}

func initDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	variables := map[string]string{
		"DB_HOST":     host,
		"DB_PORT":     port,
		"DB_USER":     user,
		"DB_PASSWORD": password,
		"DB_NAME":     dbname,
		"DB_SSLMODE":  sslmode,
	}

	for key, value := range variables {
		if value == "" {
			log.Fatalf("Environment variable %s is not set", key)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	if err = executeMigration(db); err != nil {
		log.Fatal("Error during migrating scripts", err)
	}

	log.Println("Successfully connected to the database")
	return db
}

func executeMigration(db *sql.DB) error {
	migrationPath := "db/migrations"

	if err := goose.Up(db, migrationPath); err != nil {
		return err
	}

	return nil
}
