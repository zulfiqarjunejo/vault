package assets

import (
	"time"
)

type Asset struct {
	CreatedAt time.Time   `json:"createdAt" bson:"created_at"`
	Id        string      `json:"id" bson:"id"`
	Name      string      `json:"name" bson:"name"`
	Type      string      `json:"type" bson:"type"`
	Files     []AssetFile `json:"files" bson:"files"`
}

type AssetFile struct {
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	Id        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
}
