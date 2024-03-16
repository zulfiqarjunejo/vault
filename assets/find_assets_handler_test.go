package assets

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func TestFindAssets(t *testing.T) {

	ctx := context.Background()

	mongodbContainer, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	mongoUrl, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		t.Errorf("failed to get connection string: %s", err)
	}

	t.Run("should respond back with status 200", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mongoClient, err := mongo.Connect(ctx, mongoOptions.Client().ApplyURI(mongoUrl))
		if err != nil {
			log.Fatalf("MongoDB connection failed: %+v", err.Error())
		}
		defer func() {
			if err = mongoClient.Disconnect(context.Background()); err != nil {
				panic(err)
			}
		}()

		service := NewAssetService(mongoClient)
		handler := NewFindAssetsHandler(&service)

		request, _ := http.NewRequest("GET", "/assets", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("response code is wrong, got %d want %d \n", got, want)
		}
	})
}
