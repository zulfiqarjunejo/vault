package assets

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/zulfiqarjunejo/vault/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type AssetIntegrationSuite struct {
	suite.Suite
	mongo   *mongo.Client
	model   AssetModel
	handler *AssetHandler
}

func (s *AssetIntegrationSuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbContainer, _ := mongodb.RunContainer(ctx)
	mongoUrl, _ := mongodbContainer.ConnectionString(ctx)
	mongoClient, _ := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(mongoUrl))

	model := NewMongoAssetModel(mongoClient)
	handler := NewAssetHandler(model)

	s.mongo = mongoClient
	s.model = model
	s.handler = handler
}

func TestAssetIntegrationSuite(t *testing.T) {
	suite.Run(t, new(AssetIntegrationSuite))
}

func (s *AssetIntegrationSuite) TestHandleCreateAsset() {
	s.Run("should respond with status 201", func() {
		b, _ := json.Marshal(Asset{
			Name: "test asset",
			Type: "credentials",
		})

		request, _ := http.NewRequest(http.MethodPost, "/assets", bytes.NewReader(b))
		response := httptest.NewRecorder()

		s.handler.HandleCreateAsset(response, request)

		s.Assert().Equal(http.StatusCreated, response.Result().StatusCode)

		var asset Asset
		json.NewDecoder(response.Body).Decode(&asset)

		s.Assert().Equal("test asset", asset.Name)
		s.Assert().Equal("credentials", asset.Type)
	})
}

func (s *AssetIntegrationSuite) TestHandleFindAssets() {
	s.Run("should respond with status 200", func() {
		request, _ := http.NewRequest(http.MethodGet, "/assets", nil)
		response := httptest.NewRecorder()

		s.handler.HandleFindAssets(response, request)

		s.Assert().Equal(http.StatusOK, response.Result().StatusCode)
	})
}

func (s *AssetIntegrationSuite) TestHandleGetAsset() {
	s.Run("should respond with status 400", func() {
		s.HTTPStatusCode(s.handler.HandleGetAsset, http.MethodGet, "/assets", nil, http.StatusBadRequest)
	})

	s.Run("should respond with status 404", func() {
		request, _ := http.NewRequest(http.MethodGet, "/assets/{id}", nil)
		request.SetPathValue("id", "some-id")
		response := httptest.NewRecorder()

		s.handler.HandleGetAsset(response, request)

		s.Assert().Equal(http.StatusNotFound, response.Result().StatusCode)
	})

	s.Run("should respond with status 200", func() {
		expected, _ := createAndReturnAsset(s.mongo)

		request, _ := http.NewRequest(http.MethodGet, "/assets/"+expected.Id, nil)
		request.SetPathValue("id", expected.Id)
		response := httptest.NewRecorder()

		s.handler.HandleGetAsset(response, request)

		var asset Asset
		json.NewDecoder(response.Body).Decode(&asset)

		s.Assert().Equal(http.StatusOK, response.Result().StatusCode)
		s.Assert().Equal(expected.Name, asset.Name)
		s.Assert().Equal(expected.Type, asset.Type)
		s.Assert().Equal(expected.Id, asset.Id)
	})
}

func (s *AssetIntegrationSuite) TestHandleGetAssetWithCORS() {
	s.Run("should allow requests from everywhere", func() {
		expected, _ := createAndReturnAsset(s.mongo)

		request, _ := http.NewRequest(http.MethodGet, "/assets/"+expected.Id, nil)
		request.SetPathValue("id", expected.Id)
		response := httptest.NewRecorder()

		handler := middleware.WithCORS(http.HandlerFunc(s.handler.HandleGetAsset))
		handler.ServeHTTP(response, request)

		s.Assert().Equal("*", response.Header().Get("Access-Control-Allow-Origin"))
	})
}

func (s *AssetIntegrationSuite) TestHandleGetAssetFiles() {
	s.Run("should respond with status 400", func() {
		s.HTTPStatusCode(s.handler.HandleGetAssetFiles, http.MethodGet, "/assets/{id}/files", nil, http.StatusBadRequest)
	})

	s.Run("should respond with status 404", func() {
		request, _ := http.NewRequest(http.MethodGet, "/assets/{id}/files", nil)
		request.SetPathValue("id", "some-id")
		response := httptest.NewRecorder()

		s.handler.HandleGetAssetFiles(response, request)

		s.Assert().Equal(http.StatusNotFound, response.Result().StatusCode)
	})

	s.Run("should return list of files", func() {
		expected, _ := createAndReturnAsset(s.mongo)

		request, _ := http.NewRequest(http.MethodGet, "/assets/{id}/files", nil)
		request.SetPathValue("id", expected.Id)
		response := httptest.NewRecorder()

		s.handler.HandleGetAssetFiles(response, request)

		s.Assert().Equal(http.StatusOK, response.Result().StatusCode)
		s.Assert().Equal("application/json", response.Header().Get("Content-Type"))

		var assetFiles []AssetFile
		json.NewDecoder(response.Body).Decode(&assetFiles)

		s.Assert().Equal(2, len(assetFiles))
	})
}

func (s *AssetIntegrationSuite) TestHandleGetAssetFilesWithCORS() {
	s.Run("shoud allow requests from everywhere", func() {
		expected, _ := createAndReturnAsset(s.mongo)

		request, _ := http.NewRequest(http.MethodGet, "/assets/{id}/files", nil)
		request.SetPathValue("id", expected.Id)
		response := httptest.NewRecorder()

		handler := middleware.WithCORS(http.HandlerFunc(s.handler.HandleGetAssetFiles))
		handler.ServeHTTP(response, request)

		s.Assert().Equal(http.StatusOK, response.Result().StatusCode)
		s.Assert().Equal("application/json", response.Header().Get("Content-Type"))
		s.Assert().Equal("*", response.Header().Get("Access-Control-Allow-Origin"))

		var assetFiles []AssetFile
		json.NewDecoder(response.Body).Decode(&assetFiles)

		s.Assert().Equal(2, len(assetFiles))
	})
}

func (s *AssetIntegrationSuite) TestUploadFile() {
	s.Run("handler should exists", func() {
		request, _ := http.NewRequest(http.MethodPost, "/assets/{id}/files", nil)
		response := httptest.NewRecorder()

		s.handler.HandleUploadFile(response, request)
	})

	s.Run("should respond with status 201", func() {
		expected, _ := createAndReturnAsset(s.mongo)

		fileContent := "This is a sample file content."
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		fileWriter, _ := writer.CreateFormFile("file", "sample.txt")
		io.Copy(fileWriter, bytes.NewBufferString(fileContent))
		writer.Close()

		request, _ := http.NewRequest(http.MethodPost, "/assets/{id}/files", body)
		request.SetPathValue("id", expected.Id)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		handler := middleware.WithCORS(http.HandlerFunc(s.handler.HandleUploadFile))
		handler.ServeHTTP(response, request)

		s.Assert().Equal(http.StatusCreated, response.Result().StatusCode)
		s.Assert().Equal("application/json", response.Header().Get("Content-Type"))
		s.Assert().Equal("*", response.Header().Get("Access-Control-Allow-Origin"))

		var assetFile AssetFile
		json.NewDecoder(response.Body).Decode(&assetFile)

		s.Assert().NotEmpty(assetFile.Id)
		s.Assert().NotEmpty(assetFile.Name)
	})
}

func (s *AssetIntegrationSuite) TestUploadFileWithCORS() {
	s.Run("should respond with status 201", func() {
		expected, _ := createAndReturnAsset(s.mongo)

		fileContent := "This is a sample file content."
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		fileWriter, _ := writer.CreateFormFile("file", "sample.txt")
		io.Copy(fileWriter, bytes.NewBufferString(fileContent))
		writer.Close()

		request, _ := http.NewRequest(http.MethodPost, "/assets/{id}/files", body)
		request.SetPathValue("id", expected.Id)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		s.handler.HandleUploadFile(response, request)

		s.Assert().Equal(http.StatusCreated, response.Result().StatusCode)
	})
}

func createAndReturnAsset(mongo *mongo.Client) (Asset, error) {
	assetsCollection := mongo.Database("vault").Collection("assets")

	asset := Asset{
		CreatedAt: time.Now(),
		Id:        uuid.NewString(),
		Name:      "Amazon Web Services",
		Type:      "file",
		Files: []AssetFile{
			{
				CreatedAt: time.Now(),
				Id:        uuid.NewString(),
				Name:      "aws.json",
			},
			{
				CreatedAt: time.Now().Add(time.Minute * 10),
				Id:        uuid.NewString(),
				Name:      "aws.json",
			},
		},
	}

	assetsCollection.InsertOne(context.Background(), asset)
	return asset, nil
}
