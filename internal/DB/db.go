package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConnectionToShortCodeDB() (collection *mongo.Collection, client *mongo.Client) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://vishuchoudhary131:Avon11Mongo@cluster0.4ost8i8.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Access a specific database
	database := client.Database("chotu-go")

	// Insert data into a collection within the database
	collection = database.Collection("shortCodes")
	return
}

func DisconnectClientConnection(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func CheckForDuplicateKey(shortCode string) (exist bool, err error) {
	exist = false
	collection, client := CreateConnectionToShortCodeDB()
	defer DisconnectClientConnection(client)

	filter := bson.M{shortCode: bson.M{"$exists": true}}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatalln(err)
	}
	defer cursor.Close(context.Background())
	if cursor.Next(context.Background()) {
		exist = true
	}
	return
}

func AddShortCode(shortCode, url string) (err error) {
	collection, client := CreateConnectionToShortCodeDB()
	defer DisconnectClientConnection(client)

	_, err = collection.InsertOne(context.Background(), bson.M{shortCode: url})
	if err != nil {
		log.Fatalln(err)
	}
	return
}
