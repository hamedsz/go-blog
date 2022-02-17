package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Insert(collection string , document interface{}) (*mongo.InsertOneResult , error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return GetDB().Collection(collection).InsertOne(ctx , document)
}
