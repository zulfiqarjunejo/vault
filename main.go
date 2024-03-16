package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zulfiqarjunejo/vault/assets"
)

func main() {
	environment, err := NewEnvironment()
	if err != nil {
		log.Fatalf("Error loading environment variables: %+v", err.Error())
	}

	mongoClient, err := NewMongoClient(environment)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %+v", err.Error())
	}
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	assetService := assets.NewAssetService(mongoClient)

	mux := http.NewServeMux()

	findAssetsHandler := assets.NewFindAssetsHandler(&assetService)
	createAssetHandler := assets.NewCreateAssetHandler(&assetService)

	mux.Handle("POST /assets", createAssetHandler)
	mux.Handle("GET /assets", findAssetsHandler)

	fmt.Printf("Server listening on PORT: %s \n", environment.Port)
	err = http.ListenAndServe(environment.Port, mux)
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err.Error())
	}
}
