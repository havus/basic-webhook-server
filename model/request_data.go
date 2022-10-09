package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Request data structure
type RequestData struct {
	UUID        		string
	AccountID				int
	RawHeaders 			string
	RawQueryStrings	string
	RawBody					string
	Method					string
	IpAddress				string
	Hostname				string
	UserAgent				string
	CreatedAt   		time.Time
}

func (request_data *RequestData) MarshalBSON() ([]byte, error) {
	if request_data.CreatedAt.IsZero() {
		request_data.CreatedAt = time.Now().UTC()
	}

	// ref: https://stackoverflow.com/questions/71902455/autofill-created-at-and-updated-at-in-golang-struct-while-pushing-into-mongodb
	type my RequestData
	return bson.Marshal((*my)(request_data))

	// also we can use:
	// return bson.Marshal(*request_data)
}
