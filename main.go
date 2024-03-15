package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zulfiqarjunejo/vault/assets"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env", ".env.local")
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err.Error())
	}

	PORT := os.Getenv("PORT")
	MONGO_URL := os.Getenv("MONGO_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(MONGO_URL))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %+v", err.Error())
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Initialize models.
	assetModel := assets.NewAssetModelImpl(client)

	mux := http.NewServeMux()

	config := Config{
		Mux:        mux,
		AssetModel: &assetModel,
	}

	err = InitRoutes(config)
	if err != nil {
		log.Fatalf("Route Initialization Failed: %s\n", err.Error())
	}

	fmt.Printf("Server listening on PORT: %s \n", PORT)
	err = http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err.Error())
	}
}
