package models

import "time"

type Todo struct {
	ID          int
	Title       string
	Description string
	Status      string
	Created_at  time.Time
	UserId      int
}
