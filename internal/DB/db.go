package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConnectionToShortCodeDB() (collection *mongo.Collection, client *mongo.Client) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Get db connection from env file
	dbConnection := os.Getenv("DB_KEY")
	opts := options.Client().ApplyURI(dbConnection).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Access a specific database
	database := client.Database("chotu-go")

	// A collection within the database
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
		log.Println(err)
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
		log.Println(err)
	}
	return
}

func GetUrlByShortCode(shortCode string) (url string, err error) {
	collection, client := CreateConnectionToShortCodeDB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return "", fmt.Errorf("database connection error: %w", err)
	}
	defer DisconnectClientConnection(client)

	filter := bson.M{"shortCode": shortCode}

	var document bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No document found for shortCode: %s", shortCode)
			return "", fmt.Errorf("shortcode not found: %s", shortCode)
		}
		log.Printf("Error querying database: %v", err)
		return "", fmt.Errorf("database query error: %w", err)
	}

	if value, ok := document["url"].(string); ok {
		return value, nil
	}

	log.Printf("URL not found in document for shortCode: %s", shortCode)
	return "", fmt.Errorf("URL not found for shortcode: %s", shortCode)
}
