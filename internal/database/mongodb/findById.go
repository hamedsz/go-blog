package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func FindById(collection string , id string) (*mongo.SingleResult ,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil{
		return nil, err
	}

	return GetDB().Collection(collection).FindOne(ctx , bson.M{
		"_id": objectId,
	}) , nil
}
