package domain

import "time"

type User struct {
	ID    int
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}
