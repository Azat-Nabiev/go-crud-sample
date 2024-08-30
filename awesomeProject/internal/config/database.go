package config

import (
	"database/sql"
	_ "github.com/lib/pq"
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

	log.Println("Successfully connected to the database")
	return db
}
