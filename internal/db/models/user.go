package models

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}
