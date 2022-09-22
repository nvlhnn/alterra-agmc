package database

import (
	"alterra-agmc/day-6/pkg/utils"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewDatabaseConfigMongo() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(utils.GoDotEnvVariable("MONGO_URI")))
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
	db := client.Database(utils.GoDotEnvVariable("MONGO_DBNAME"))


	// seeding books
	results, err :=  db.Collection("book").Find(ctx, bson.M{})
	if results.RemainingBatchLength() > 0 {
		return db
	}
	fmt.Println("Seeding")
	books := []interface{}{
        bson.D{{"name", "book 1"}, {"author", "author 1"}},
        bson.D{{"name", "book 2"}, {"author", "author 2"}},
        bson.D{{"name", "book 3"}, {"author", "author 3"}},
	}
	// insert the bson object slice using InsertMany()
	_, err = db.Collection("book").InsertMany(context.TODO(), books)
	if err != nil {
			panic(err)
	}



	return db
}
