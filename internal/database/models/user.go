package models

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
