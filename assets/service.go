package assets

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateAssetService interface {
	CreateAsset(Asset) (Asset, error)
}

type FindAssetsService interface {
	FindAssets() ([]Asset, error)
}

type AssetService struct {
	Mongo *mongo.Client
}

func NewAssetService(mongo *mongo.Client) AssetService {
	return AssetService{
		Mongo: mongo,
	}
}

func (model *AssetService) CreateAsset(asset Asset) (Asset, error) {
	asset.Id = uuid.NewString()
	asset.CreatedAt = time.Now()

	assetsCollection := model.Mongo.Database("vault").Collection("assets")
	_, err := assetsCollection.InsertOne(context.Background(), asset) // TODO: Process result

	return asset, err
}

func (model *AssetService) FindAssets() ([]Asset, error) {
	assetsCollection := model.Mongo.Database("vault").Collection("assets")

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := assetsCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	var assets []Asset
	err = cursor.All(context.TODO(), &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
