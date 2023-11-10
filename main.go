package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

var client *mongo.Client

func setupMongoDBClient() *mongo.Client {
	mongoURL := viper.GetString("mongo.url") // Read MongoDB URL from configuration
	clientOptions := options.Client().ApplyURI(mongoURL)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Unable to connect to MongoDB:", err)
	}

	return client
}

func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database(viper.GetString("mongo.database")).Collection(viper.GetString("mongo.collection"))
	_, err := collection.InsertOne(ctx, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// Other CRUD handlers...

func main() {
	// Configure Viper
	viper.SetConfigName("config") // Use a configuration file named "config" (config.yaml, config.json, etc.)
	viper.SetConfigType("yaml")   // Set the configuration file type (yaml, json, etc.)
	viper.AddConfigPath(".")      // Search for the configuration file in the current directory

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file:", err)
	}

	// Set default values for configuration parameters
	viper.SetDefault("mongo.url", "mongodb://localhost:27017")
	viper.SetDefault("mongo.database", "sri")
	viper.SetDefault("mongo.collection", "sric")

	client = setupMongoDBClient()
	defer client.Disconnect(context.Background())

	router := gin.Default()

	router.POST("/items", createItem)
	// Add routes for other CRUD operations...

	router.Run(":8888")
}
