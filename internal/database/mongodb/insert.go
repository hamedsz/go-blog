package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (db *Database) Insert(table string , data interface{}) (*string , error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := GetDB().Collection(table).InsertOne(ctx , data)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return &id , nil

	return nil, nil
}
