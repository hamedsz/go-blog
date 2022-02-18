package mongodb

import (
	"context"
	"golang-blog/database/base"
	"log"
)

func (db *Database) Find(query base.Query) ([]map[string]interface{}, error) {
	cursor, err := GetDB().Collection(db.Model.Name()).Find(context.TODO(), query.GetFilter())

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