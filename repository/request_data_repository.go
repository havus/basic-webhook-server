package repository

import (
	"context"
	"log"

	"github.com/havus/go-webhook-server/model/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ---------- Interface ----------
type RequestDataRepository interface {
	Insert(ctx context.Context, request entity.RequestData) (entity.RequestData, error)
	FindAll(ctx context.Context, accountId string) ([]entity.RequestData, error)
}

// ---------- Interface Implementation ----------
type RequestDataRepositoryImpl struct {
	db *mongo.Database
}

func NewRequestDataRepository(db *mongo.Database) *RequestDataRepositoryImpl {
	return &RequestDataRepositoryImpl{
		db: db,
	}
}

func (repository *RequestDataRepositoryImpl) Insert(ctx context.Context, request_data entity.RequestData) (entity.RequestData, error) {
	request_data_marshalled, err := bson.Marshal(&request_data)
	if err != nil {
		return entity.RequestData{}, err
	}

	if _, err := repository.db.Collection("requests").InsertOne(ctx, request_data_marshalled); err != nil {
		return entity.RequestData{}, err
	}

	return request_data, nil
}

func (repository *RequestDataRepositoryImpl) FindAll(ctx context.Context, accountId string) ([]entity.RequestData, error) {
	var requests []entity.RequestData
	filter := bson.D{
		primitive.E{
			Key: "account_id",
			Value: accountId,
		},
	}

	cursor, err := repository.db.Collection("requests").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var row entity.RequestData

		err := cursor.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		requests = append(requests, row)
	}

	return requests, nil
}