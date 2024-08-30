package models

import "time"

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Age     int       `json:"age"`
	Since   time.Time `json:"since"`
	Books   []Book    `json:"books"`
}
