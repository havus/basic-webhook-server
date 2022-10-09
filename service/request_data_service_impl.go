package service

import (
	"basic-webhook-server/handler/response"
	"basic-webhook-server/model"
	"basic-webhook-server/repository"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (service *RequestDataServiceImpl) Create(ctx *gin.Context, request_method string) (*response.RequestDataResponse, error) {
	// see: byte to string convertion https://stackoverflow.com/questions/40632802/how-to-convert-byte-array-to-string-in-go
	uuid := uuid.New().String()

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	header_marshalled, err := json.Marshal(ctx.Request.Header)
	if err != nil {
		return nil, err
	}

	query_string_marshalled, err := json.Marshal(ctx.Request.URL.Query())
	if err != nil {
		return nil, err
	}

	request_data, err := service.requestDataRepository.Insert(
		ctx,
		model.RequestData{
			UUID:            uuid,
			AccountID:       1,
			RawHeaders:      string(header_marshalled[:]),
			RawQueryStrings: string(query_string_marshalled[:]),
			RawBody:         string(jsonData[:]),
			Method:          request_method,
			IpAddress:       ctx.ClientIP(),
			Hostname:        ctx.Request.Host,
			UserAgent:       ctx.GetHeader("User-Agent"),
		},
	)

	if err != nil {
		return nil, err
	}

	return &response.RequestDataResponse{
		UUID:            request_data.UUID,
		AccountID:       request_data.AccountID,
		RawHeaders:      request_data.RawHeaders,
		RawQueryStrings: request_data.RawQueryStrings,
		RawBody:         request_data.RawBody,
		Method:          request_data.Method,
		IpAddress:       request_data.IpAddress,
		Hostname:        request_data.Hostname,
		UserAgent:       request_data.UserAgent,
		CreatedAt:       request_data.CreatedAt.String(),
	}, nil
}
