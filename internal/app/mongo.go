package app

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"syscall"
)

// NewMongoClient creates a new Mongo client
func (appCfg *AppConfig) NewMongoClient() (*mongo.Client, error) {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(appCfg.MongoConfig.URI))
	if err != nil {
		return nil, err
	}
	return mongoClient, err
}

// MongoConnect establishes connection with and pings the MongoDB server
func (appCfg *AppConfig) MongoConnect(ctx context.Context, stopChan chan os.Signal) {
	if err := appCfg.MongoConfig.MongoClient.Connect(ctx); err != nil {
		log.Errorf("error connecting to mongoDB server: %s", err.Error())
		stopChan <- syscall.SIGTERM
	} else if err := appCfg.MongoConfig.MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Errorf("error pinging mongoDB server: %s", err.Error())
		stopChan <- syscall.SIGTERM
	} else {
		log.Info("Connected to MongoDB server...")
	}
}
