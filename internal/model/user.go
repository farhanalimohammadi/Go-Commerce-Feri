package model

import "time"

type User struct {
	ID int64
	Email string
	PsswordHash string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}