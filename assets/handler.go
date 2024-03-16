package assets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AssetHandler struct {
	assetModel AssetModel
}

func NewAssetHandler(assetModel AssetModel) *AssetHandler {
	return &AssetHandler{
		assetModel,
	}
}

func (h *AssetHandler) FindAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.assetModel.FindAssets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected Error: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&assets)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}

func (h *AssetHandler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	var asset Asset

	err := json.NewDecoder(r.Body).Decode(&asset)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	asset, err = h.assetModel.CreateAsset(asset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected Error: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&asset)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}
