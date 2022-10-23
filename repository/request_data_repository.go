package repository

import (
	"context"
	"fmt"

	"github.com/havus/go-webhook-server/model/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestDataRepository interface {
	Insert(ctx context.Context, request entity.RequestData) (*entity.RequestData, error)
}

type RequestDataRepositoryImpl struct {
	db *mongo.Database
}

func NewRequestDataRepository(db *mongo.Database) *RequestDataRepositoryImpl {
	return &RequestDataRepositoryImpl{
		db: db,
	}
}

func (repository *RequestDataRepositoryImpl) Insert(ctx context.Context, request_data entity.RequestData) (*entity.RequestData, error) {
	request_data_marshalled, err := bson.Marshal(&request_data)
	if err != nil {
		return nil, err
	}

	if _, err := repository.db.Collection("requests").InsertOne(ctx, request_data_marshalled); err != nil {
		return nil, err
	}

	fmt.Printf("Insert success uuid: %s\n", request_data.UUID)

	return &request_data, nil
}
