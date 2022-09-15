package repository

import (
	"basic-webhook-server/model"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type RequestDataRepositoryImpl struct {
}

func NewRequestDataRepository() *RequestDataRepositoryImpl {
	return &RequestDataRepositoryImpl{}
}

func (repository *RequestDataRepositoryImpl) Set(ctx context.Context, request_data model.RequestData) model.RequestData {
	// panic("asd")
	uuid := uuid.New().String()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, uuid, request_data.RawBody, 180*time.Second).Err()

	if err != nil {
		panic(err)
	}

	request_data.UUID = uuid

	return request_data
}
