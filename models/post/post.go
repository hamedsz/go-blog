package post

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-blog/database/base"
	"golang-blog/database/mongodb"
	"time"
)

type Model struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string              `bson:"title,omitempty" json:"title"`
	Body      string              `bson:"body,omitempty" json:"body"`
	CreatedAt time.Time           `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at,omitempty" json:"updated_at"`
}

func (model *Model) Name() string{
	return "posts"
}

func (model *Model) GetDB() base.Database {
	db := mongodb.Database{
		Model: model,
	}
	return base.GetDatabase(&db)
}

func (model *Model) SetCreatedAt(createdAt time.Time) {
	model.CreatedAt = createdAt
}

func (model *Model) SetUpdatedAt(updatedAt time.Time) {
	model.UpdatedAt = updatedAt
}

func List(data []map[string]interface{}) ([]Model, error){

	var result []Model

	for _, element := range data {

		data, err := bson.Marshal(element)

		if err != nil{
			return nil, err
		}

		var model Model

		err = bson.Unmarshal(data , &model)

		if err != nil{
			return nil, err
		}

		result = append(result, model)
	}

	return result , nil
}