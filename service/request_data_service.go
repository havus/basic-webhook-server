package service

import (
	"encoding/json"
	"io"

	"github.com/havus/go-webhook-server/helper"
	"github.com/havus/go-webhook-server/model/entity"
	"github.com/havus/go-webhook-server/model/http"
	"github.com/havus/go-webhook-server/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestDataService interface {
	Create(ctx *gin.Context, requestMethod string) (http.RequestDataResponse, error)
	GetAllByAccountId(ctx *gin.Context, minId interface{}, maxId interface{}) ([]http.RequestDataResponse, error)
}

type RequestDataServiceImpl struct {
	requestDataRepository repository.RequestDataRepository
}

func NewRequestDataService(requestDataRepository repository.RequestDataRepository) *RequestDataServiceImpl {
	return &RequestDataServiceImpl{
		requestDataRepository: requestDataRepository,
	}
}

func (service *RequestDataServiceImpl) Create(ctx *gin.Context, requestMethod string) (http.RequestDataResponse, error) {
	// Byte to string convertion. See https://stackoverflow.com/questions/40632802/how-to-convert-byte-array-to-string-in-go
	uuid 			:= uuid.New().String()
	accountId := ctx.Param("account_id")

	rawData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return http.RequestDataResponse{}, err
	}

	headerMarshalled, err := json.Marshal(ctx.Request.Header)
	if err != nil {
		return http.RequestDataResponse{}, err
	}

	queryStringMarshalled, err := json.Marshal(ctx.Request.URL.Query())
	if err != nil {
		return http.RequestDataResponse{}, err
	}

	request_data, err := service.requestDataRepository.Insert(
		ctx,
		entity.RequestData{
			UUID:							uuid,
			AccountID: 				accountId,
			RawHeaders: 			string(headerMarshalled[:]),
			RawQueryStrings: 	string(queryStringMarshalled[:]),
			RawBody: 					string(rawData[:]),
			Method: 					requestMethod,
			IpAddress: 				ctx.ClientIP(),
			Hostname: 				ctx.Request.Host,
			UserAgent: 				ctx.GetHeader("User-Agent"),
		},
	)

	if err != nil {
		return http.RequestDataResponse{}, err
	}

	return helper.ToRequestDataResponse(request_data), nil
}

func (service *RequestDataServiceImpl) GetAllByAccountId(ctx *gin.Context, minId interface{}, maxId interface{}) ([]http.RequestDataResponse, error) {
	accountId 				:= ctx.Param("account_id")
	requestDatas, err := service.requestDataRepository.FindAll(ctx, accountId, minId, maxId)

	if err != nil {
		return nil, err
	}

	return helper.ToRequestDataResponses(requestDatas), nil
}
