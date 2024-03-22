package assets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/zulfiqarjunejo/vault/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssetHandler struct {
	assetModel AssetModel
}

func NewAssetHandler(assetModel AssetModel) *AssetHandler {
	return &AssetHandler{
		assetModel,
	}
}

func (h *AssetHandler) HandleGetAsset(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id is required!"))
		return
	}

	asset, err := h.assetModel.GetAssetById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("asset not found"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected Error: " + err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(&asset)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}

func (h *AssetHandler) HandleFindAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.assetModel.FindAssets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected Error: " + err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(&assets)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}

func (h *AssetHandler) HandleCreateAsset(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&asset)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}

func (h *AssetHandler) HandleGetAssetFiles(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id is required!"))
		return
	}

	files, err := h.assetModel.GetAssetFiles(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("asset not found"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected Error: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&files)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}

func (h *AssetHandler) HandleUploadFile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	dir, err := utils.CreateUploadDirectoryFor(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Missing or malformed file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a UUID for the file
	uuid := uuid.New().String()

	// Get the extension of the original filename
	ext := filepath.Ext(fileHeader.Filename)

	// Construct the new filename with UUID
	newFilename := fmt.Sprintf("%s_%s%s", strings.TrimSuffix(fileHeader.Filename, ext), uuid, ext)

	f, err := os.Create(filepath.Join(dir, newFilename))
	if err != nil {
		http.Error(w, "Unable to create file on server", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Unable to copy file content", http.StatusInternalServerError)
		return
	}

	assetFile, err := h.assetModel.CreateAssetFile(id, newFilename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected Error: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&assetFile)
	if err != nil {
		fmt.Printf("Json Encoding Error: %v \n", err.Error())
	}
}
