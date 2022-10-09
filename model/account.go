package model

import "time"

// Account data structure
type Account struct {
	ID           int
	Email        string
	UniqueUrl    string
	StatusCode   int
	ResponseBody string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
