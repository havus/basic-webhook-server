package model

import "time"

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
