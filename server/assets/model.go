package assets

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AssetModel interface {
	CreateAsset(Asset) (Asset, error)
	FindAssets() ([]Asset, error)
	GetAssetById(string) (*Asset, error)
	GetAssetFiles(string) ([]AssetFile, error)
	CreateAssetFile(string, string) (*AssetFile, error)
}

type MongoAssetModel struct {
	mongo *mongo.Client
}

func NewMongoAssetModel(mongo *mongo.Client) *MongoAssetModel {
	return &MongoAssetModel{
		mongo,
	}
}

func (m *MongoAssetModel) CreateAsset(asset Asset) (Asset, error) {
	asset.Id = uuid.NewString()
	asset.CreatedAt = time.Now()

	assetsCollection := m.mongo.Database("vault").Collection("assets")
	_, err := assetsCollection.InsertOne(context.Background(), asset) // TODO: Process result

	return asset, err
}

func (m *MongoAssetModel) FindAssets() ([]Asset, error) {
	assetsCollection := m.mongo.Database("vault").Collection("assets")

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

func (m *MongoAssetModel) GetAssetById(id string) (*Asset, error) {
	assetsCollection := m.mongo.Database("vault").Collection("assets")

	var asset Asset

	filter := bson.D{
		{Key: "id", Value: id},
	}

	err := assetsCollection.FindOne(context.Background(), filter).Decode(&asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (m *MongoAssetModel) GetAssetFiles(id string) ([]AssetFile, error) {
	asset, err := m.GetAssetById(id)
	if err != nil {
		return nil, err
	}

	return asset.Files, nil
}

func (m *MongoAssetModel) CreateAssetFile(id string, name string) (*AssetFile, error) {
	asset, err := m.GetAssetById(id)
	if err != nil {
		return nil, err
	}

	newAssetFile := AssetFile{
		CreatedAt: time.Now(),
		Id:        uuid.NewString(),
		Name:      name,
	}

	if asset.Files == nil {
		asset.Files = []AssetFile{newAssetFile}
	} else {
		asset.Files = append(asset.Files, newAssetFile)
	}

	assetsCollection := m.mongo.Database("vault").Collection("assets")

	filter := bson.D{
		{Key: "id", Value: id},
	}
	update := bson.M{
		"$set": bson.D{
			{Key: "files", Value: asset.Files},
		},
	}

	_, err = assetsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &newAssetFile, nil
}
