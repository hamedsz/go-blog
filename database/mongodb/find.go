package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/hamedsz/go-blog/database/base"
	"log"
)

func (db *Database) Find(query base.Query) ([]map[string]interface{}, error) {

	option := options.Find()

	option.SetLimit(query.GetLimit())
	option.SetSkip(query.GetSkip())

	cursor, err := GetDB().Collection(db.Model.Name()).Find(context.TODO(), query.GetFilter(), option)

	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem map[string]interface{}
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results =append(results, elem)
	}

	return results, nil
}