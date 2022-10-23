package entity

import "time"

// Account data structure
type Account struct {
	ID        		string
	Email					string
	UniqueUrl			string
	StatusCode		int
	ResponseBody	string
	CreatedAt   	time.Time
	UpdatedAt 		time.Time
}
