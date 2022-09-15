package service

import (
	"basic-webhook-server/handler/request"
	"basic-webhook-server/handler/response"
	"context"
)

type RequestDataService interface {
	Create(ctx context.Context, request request.RequestDataRequest) response.RequestDataResponse
}
