package post

import (
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-blog/database/base"
	"golang-blog/database/mongodb"
)

type Model struct {
	ID    *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string              `bson:"title,omitempty" json:"title"`
	Body  string              `bson:"body,omitempty" json:"body"`
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

func List(data []map[string]interface{}) ([]Model, error){

	var result []Model

	for _, element := range data {

		element["ID"] = element["_id"]

		var model Model

		err := mapstructure.Decode(element , &model)

		if err != nil{
			return nil, err
		}

		result = append(result, model)
	}

	return result , nil
}