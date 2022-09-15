package repository

import (
	"basic-webhook-server/model"
	"context"
)


type RequestDataRepository interface {
	Set(ctx context.Context, request model.RequestData) model.RequestData
}
