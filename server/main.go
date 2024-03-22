package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zulfiqarjunejo/vault/assets"
	"github.com/zulfiqarjunejo/vault/middleware"
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

	assetModel := assets.NewMongoAssetModel(mongoClient)

	mux := http.NewServeMux()

	assetHandler := assets.NewAssetHandler(assetModel)

	mux.HandleFunc("GET /assets", assetHandler.HandleFindAssets)
	mux.HandleFunc("GET /assets/{id}", assetHandler.HandleGetAsset)
	mux.HandleFunc("GET /assets/{id}/files", assetHandler.HandleGetAssetFiles)

	mux.HandleFunc("POST /assets/{id}/files", assetHandler.HandleUploadFile)
	mux.HandleFunc("POST /assets", assetHandler.HandleCreateAsset)

	fmt.Printf("Server listening on PORT: %s\n", environment.Port)

	s := http.Server{
		Addr:         fmt.Sprintf(":%s", environment.Port),
		Handler:      middleware.WithCORS(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
