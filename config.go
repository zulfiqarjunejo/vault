package main

import (
	"net/http"

	"github.com/zulfiqarjunejo/vault/assets"
)

type Config struct {
	Mux        *http.ServeMux
	AssetModel assets.AssetModel
}
