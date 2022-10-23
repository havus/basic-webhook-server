package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddFilterGreaterThan(minId interface{}, filter *bson.D) {
	if minId == nil {
		return
	}

	objID, err := primitive.ObjectIDFromHex(minId.(string))
	if err != nil {
		return
	}

	*filter = append(
		*filter,
		primitive.E{
			Key: "_id",
			Value: bson.D{
				primitive.E{
					Key: "$gt",
					Value: objID,
				},
			},
		},
	)
}

func AddFilterLessThan(maxId interface{}, filter *bson.D) {
	if maxId == nil {
		return
	}

	objID, err := primitive.ObjectIDFromHex(maxId.(string))
	if err != nil {
		return
	}

	*filter = append(
		*filter,
		primitive.E{
			Key: "_id",
			Value: bson.D{
				primitive.E{
					Key: "$lt",
					Value: objID,
				},
			},
		},
	)
}
