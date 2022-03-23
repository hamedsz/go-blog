package domain

import (
	"github.com/hamedsz/go-blog/internal/database/base"
	"github.com/hamedsz/go-blog/internal/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (model *Post) Name() string{
	return "posts"
}

func (model *Post) GetDB() base.Database {
	db := mongodb.Database{
		Model: model,
	}
	return base.GetDatabase(&db)
}

func (model *Post) SetCreatedAt(createdAt time.Time) {
	model.CreatedAt = createdAt
}

func (model *Post) SetUpdatedAt(updatedAt time.Time) {
	model.UpdatedAt = updatedAt
}

func List(data []map[string]interface{}) ([]Post, error){

	var result []Post

	for _, element := range data {

		data, err := bson.Marshal(element)

		if err != nil{
			return nil, err
		}

		var model Post

		err = bson.Unmarshal(data , &model)

		if err != nil{
			return nil, err
		}

		result = append(result, model)
	}

	return result , nil
}
