package repository

import (
	"basic-webhook-server/model"
	"context"
)


type RequestDataRepository interface {
	Insert(ctx context.Context, request model.RequestData) (*model.RequestData, error)
}
