package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client  *mongo.Client
	Bullets *mongo.Collection
	Cases   *mongo.Collection
	Powders *mongo.Collection
	Primers *mongo.Collection
)

func Init(db string) error {
	opts := options.Client().ApplyURI("mongodb://localhost:27017") // local db

	localClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	Client = localClient

	Bullets = Client.Database(db).Collection("bullets")
	Cases = Client.Database(db).Collection("cases")
	Powders = Client.Database(db).Collection("powders")
	Primers = Client.Database(db).Collection("primers")

	err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	return err

}

func Close() error {
	return Client.Disconnect(context.Background())
}
