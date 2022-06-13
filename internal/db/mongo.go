package db

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoClient creates a new Mongo client
func NewMongoClient() (*mongo.Client, error) {
	mongoURI := flag.String("mongo-uri", "mongodb://development:testpassword@localhost:27017", "URI for establishing connection to mongoDB") // TODO: update mongoDB URI via cmd line arg
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(*mongoURI))
	if err != nil {
		return nil, err
	}
	return mongoClient, err
}

// DBConnect establishes connection with and pings the database server
func (dbConfig *config) DBConnect(ctx context.Context) error {
	if err := dbConfig.mongoClient.Connect(ctx); err != nil {
		log.Errorf("error connecting to mongoDB server: %s", err.Error())
		return err
	}
	if err := dbConfig.mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Errorf("error pinging mongoDB server: %s", err.Error())
		return err
	}
	log.Info("Connected to MongoDB server...")
	return nil
}
