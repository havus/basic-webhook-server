package http

type RequestDataResponse struct {
	ID							string	`json:"id"`
	UUID        		string 	`json:"uuid"`
	AccountID				string	`json:"account_id"`
	RawHeaders 			string 	`json:"raw_headers"`
	RawQueryStrings	string 	`json:"raw_query_strings"`
	RawBody					string 	`json:"raw_body"`
	Method					string 	`json:"method"`
	IpAddress				string 	`json:"ip_address"`
	Hostname				string 	`json:"hostname"`
	UserAgent				string 	`json:"user_agent"`
	CreatedAt   		string 	`json:"created_at"`
}
