package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-blog/config/database"
	"log"
	"time"
)

func connectDB() *mongo.Client  {
	client, err := mongo.NewClient(options.Client().ApplyURI(database.GetUri()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

//Client instance
var db *mongo.Client;

func GetDB() *mongo.Database {
	if db == nil {
		db = connectDB()
	}

	return db.Database(database.GetDatabaseName())
}