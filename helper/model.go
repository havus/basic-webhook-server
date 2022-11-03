package helper

import (
	"github.com/havus/go-webhook-server/model/entity"
	"github.com/havus/go-webhook-server/model/http"
)

func ToRequestDataResponse(requestData entity.RequestData) http.RequestDataResponse {
	return http.RequestDataResponse{
		ID: 							requestData.ID,
		UUID: 						requestData.UUID,
		Url: 							requestData.Url,
		AccountID: 				requestData.AccountID,
		RawHeaders: 			requestData.RawHeaders,
		RawQueryStrings: 	requestData.RawQueryStrings,
		RawBody: 					requestData.RawBody,
		Method: 					requestData.Method,
		IpAddress: 				requestData.IpAddress,
		Hostname: 				requestData.Hostname,
		UserAgent: 				requestData.UserAgent,
		CreatedAt: 				requestData.CreatedAt.String(),
	}
}

func ToRequestDataResponses(requestDatas []entity.RequestData) []http.RequestDataResponse {
	var requestDataResponses []http.RequestDataResponse

	for _, requestData := range requestDatas {
		requestDataResponses = append(requestDataResponses, ToRequestDataResponse(requestData))
	}

	return requestDataResponses
}
