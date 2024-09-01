package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func InitDB() *sql.DB {
	dsn := "host=localhost port=5436 user=anabiev password=qwerty007 dbname=awesome sslmode=disable"

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
