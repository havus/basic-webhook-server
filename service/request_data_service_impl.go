package service

import (
	"basic-webhook-server/handler/request"
	"basic-webhook-server/handler/response"
	"basic-webhook-server/model"
	"basic-webhook-server/repository"
	"context"
)

type RequestDataServiceImpl struct {
	requestDataRepository repository.RequestDataRepository
	// DB *redis.Client
}

func NewRequestDataService(requestDataRepository repository.RequestDataRepository) *RequestDataServiceImpl {
	return &RequestDataServiceImpl{
		requestDataRepository: requestDataRepository,
	}
}

func (service *RequestDataServiceImpl) Create(ctx context.Context, request request.RequestDataRequest) (response.RequestDataResponse) {
	request_data := service.requestDataRepository.Set(
		ctx,
		model.RequestData{
			AccountID: 				request.AccountID,
			RawHeaders: 			request.RawHeaders,
			RawQueryStrings: 	request.RawQueryStrings,
			RawBody: 					request.RawBody,
			Method: 					request.Method,
			IpAddress: 				request.IpAddress,
			Hostname: 				request.Hostname,
			UserAgent: 				request.UserAgent,
		},
	)

	return response.RequestDataResponse{
		UUID: 						request_data.UUID,
		AccountID: 				request_data.AccountID,
		RawHeaders: 			request_data.RawHeaders,
		RawQueryStrings: 	request_data.RawQueryStrings,
		RawBody: 					request_data.RawBody,
		Method: 					request_data.Method,
		IpAddress: 				request_data.IpAddress,
		Hostname: 				request_data.Hostname,
		UserAgent: 				request_data.UserAgent,
	}
}
