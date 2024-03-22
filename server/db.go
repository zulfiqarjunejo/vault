package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(environment Environment) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return mongo.Connect(ctx, mongoOptions.Client().ApplyURI(environment.MongoUrl))
}
