package MDB

import (
	"context"
	"fmt"
	"lyked-backend/internal/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var mongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectMongo(dbName string) (*mongo.Database, error) {
	utils.LoadEnv()
	mongoURI := utils.GetEnv("MONGO_URI", "BACKUP_MONGO_URI")
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	mongoClient = client
	MongoDB = mongoClient.Database(dbName)
	if MongoDB == nil {
		return nil, fmt.Errorf("failed to connect to database '%s'", dbName)
	}
	fmt.Print("Connected to database: ", dbName, "\n")

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return MongoDB, nil
}

func GetCollection(CName string) (*mongo.Collection, error) {
	if mongoClient == nil {
		return nil, fmt.Errorf("MongoDB client is not connected")
	}

	Collection := MongoDB.Collection(CName)
	if Collection == nil {
		return nil, fmt.Errorf("failed to connect to collection '%s'", CName)
	}
	fmt.Print("Connected to collection: ", CName, "\n")

	return Collection, nil
}
