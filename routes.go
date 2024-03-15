package main

import (
	"net/http"

	"github.com/zulfiqarjunejo/vault/assets"
)

func InitRoutes(config Config) error {
	// assets
	findAssetsHandler := http.Handler(assets.NewFindAssetsHandler(config.AssetModel))
	createAssetHandler := http.Handler(assets.NewCreateAssetHandler(config.AssetModel))

	config.Mux.Handle("POST /assets", createAssetHandler)
	config.Mux.Handle("GET /assets", findAssetsHandler)

	return nil
}
