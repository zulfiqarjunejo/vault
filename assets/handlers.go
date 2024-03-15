package assets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewCreateAssetHandler(model AssetModel) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var asset Asset

		err := json.NewDecoder(r.Body).Decode(&asset)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		asset, err = model.CreateAsset(asset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unexpected Error: " + err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&asset)
		if err != nil {
			fmt.Printf("Json Encoding Error: %v \n", err.Error())
		}
	}

	return http.HandlerFunc(fn)
}

func NewFindAssetsHandler(model AssetModel) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		assets, err := model.FindAssets()
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

	return http.HandlerFunc(fn)
}
