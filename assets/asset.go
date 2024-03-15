package assets

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	Id        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Type      string    `json:"type" bson:"type"`
}

type AssetModel interface {
	CreateAsset(Asset) (Asset, error)
	FindAssets() ([]Asset, error)
}

type NoOpAssetModel struct{}

func (model NoOpAssetModel) CreateAsset(asset Asset) (Asset, error) {
	return asset, nil
}

func (model NoOpAssetModel) FindAssets() ([]Asset, error) {
	assets := []Asset{
		{
			Id:   uuid.NewString(),
			Name: "AWS",
			Type: "credentials",
		},
		{
			Id:   uuid.NewString(),
			Name: "Twitter",
			Type: "credentials",
		},
	}

	return assets, nil
}
