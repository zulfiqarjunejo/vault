package assets

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type AssetIntegrationSuite struct {
	suite.Suite
	mongo *mongo.Client
}

func (s *AssetIntegrationSuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbContainer, _ := mongodb.RunContainer(ctx)
	mongoUrl, _ := mongodbContainer.ConnectionString(ctx)
	mongoClient, _ := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(mongoUrl))

	s.mongo = mongoClient
}

func TestAssetIntegrationSuite(t *testing.T) {
	suite.Run(t, new(AssetIntegrationSuite))
}

func (s *AssetIntegrationSuite) TestCreateAsset() {
	model := NewMongoAssetModel(s.mongo)
	handler := NewAssetHandler(model)

	b, _ := json.Marshal(Asset{
		Name: "test asset",
		Type: "credentials",
	})

	request, _ := http.NewRequest(http.MethodPost, "/assets", bytes.NewReader(b))
	response := httptest.NewRecorder()

	handler.CreateAsset(response, request)

	s.Assert().Equal(http.StatusCreated, response.Result().StatusCode)

	var asset Asset
	json.NewDecoder(response.Body).Decode(&asset)

	s.Assert().Equal("test asset", asset.Name)
	s.Assert().Equal("credentials", asset.Type)
}

func (s *AssetIntegrationSuite) TestFindAssets() {
	model := NewMongoAssetModel(s.mongo)
	handler := NewAssetHandler(model)

	request, _ := http.NewRequest(http.MethodGet, "/assets", nil)
	response := httptest.NewRecorder()

	handler.FindAssets(response, request)

	s.Assert().Equal(http.StatusOK, response.Result().StatusCode)
}
