package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBName                     = "db2025"
	MahasiswaCollection        = "data_mahasiswa"
	MongoString         string = os.Getenv("MONGOSTRING")
	UserCollection             = "user"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	MongoString = os.Getenv("MONGOSTRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}
