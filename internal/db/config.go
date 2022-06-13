package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// Config for db operations
type Config interface { // we are exposing db.Config interface & its any required function outside the current package
	DBConnect(context.Context) error
}

type config struct {
	mongoClient *mongo.Client
}

// New creates new mongo client
func New(mongoClient *mongo.Client) Config {
	return &config{
		mongoClient: mongoClient,
	}
}
