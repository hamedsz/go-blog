package mongodb

import (
	"github.com/fatih/structs"
	"github.com/hamedsz/go-blog/internal/database/base"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (db *Database) Save() error{
	mapData := structs.Map(db.Model)

	if mapData["Id"] == nil{
		return insert(db.Model, db)
	}

	return update(db.Model, db)
}

func insert(model base.Model, db *Database) error{
	model.SetCreatedAt(time.Now())
	model.SetUpdatedAt(time.Now())

	mapData := structs.Map(model)

	result, err := db.Insert(model.Name(), model)

	if err != nil{
		return err
	}

	mapData["ID"], err = primitive.ObjectIDFromHex(*result)
	if err != nil{
		return err
	}


	return mapstructure.Decode(mapData , &model)
}

func update(model base.Model, db *Database) error {
	model.SetUpdatedAt(time.Now())

	mapData := structs.Map(model)

	return db.Update(model.Name(), bson.M{
		"_id": mapData["Id"],
	}, model);
}