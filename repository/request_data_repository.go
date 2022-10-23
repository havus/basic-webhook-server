package repository

import (
	"context"
	"log"

	"github.com/havus/go-webhook-server/helper"
	"github.com/havus/go-webhook-server/model/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ---------- Interface ----------
type RequestDataRepository interface {
	Insert(ctx context.Context, request entity.RequestData) (entity.RequestData, error)
	FindAll(ctx context.Context, accountId string, minId interface{}, maxId interface{}) ([]entity.RequestData, error)
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

	result, err := repository.db.Collection("requests").InsertOne(ctx, request_data_marshalled)
	if err != nil {
		return entity.RequestData{}, err
	}

	request_data.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return request_data, nil
}

func (repository *RequestDataRepositoryImpl) FindAll(ctx context.Context, accountId string, minId interface{}, maxId interface{}) ([]entity.RequestData, error) {
	var requests []entity.RequestData

	filter := bson.D{
		primitive.E{
			Key: "account_id",
			Value: accountId,
		},
	}

	helper.AddFilterGreaterThan(minId, &filter)
	helper.AddFilterLessThan(maxId, &filter)

	opts := options.Find().SetSort(bson.D{
		primitive.E{Key: "_id", Value: -1},
	})

	cursor, err := repository.db.Collection("requests").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

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