package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Port     string
	MongoUrl string
}

func NewEnvironment() (Environment, error) {
	envFiles := []string{".env", ".env.local"}

	err := godotenv.Load(envFiles...)
	if err != nil {
		return Environment{}, err
	}

	port := os.Getenv("PORT")
	mongoUrl := os.Getenv("MONGO_URL")

	return Environment{
		Port:     port,
		MongoUrl: mongoUrl,
	}, nil
}
