package models

import "time"

type User struct {
	ID         string
	FirstName  string
	LastName   string
	MiddleName *string
	Email      *string
	Phone      *string
	CreatedAt  time.Time
	Birthdate  time.Time
}
