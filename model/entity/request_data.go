package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Request data structure
type RequestData struct {
	ID        			string 		`bson:"_id,omitempty"`
	UUID        		string 		`bson:"uuid"`
	AccountID				string 		`bson:"account_id"`
	RawHeaders 			string 		`bson:"raw_headers"`
	RawQueryStrings	string 		`bson:"raw_query_strings"`
	RawBody					string 		`bson:"raw_body"`
	Url							string		`bson:"url"`
	Method					string 		`bson:"method"`
	IpAddress				string 		`bson:"ip_address"`
	Hostname				string 		`bson:"hostname"`
	UserAgent				string 		`bson:"user_agent"`
	CreatedAt   		time.Time	`bson:"created_at"`
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
