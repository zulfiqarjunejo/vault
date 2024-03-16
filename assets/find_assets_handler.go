package assets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewFindAssetsHandler(model FindAssetsService) http.Handler {
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
