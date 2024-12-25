package models

import "time"

type User struct {
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
