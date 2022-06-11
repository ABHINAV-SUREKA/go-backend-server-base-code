package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ServerConfig struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MongoConfig struct {
	MongoClient *mongo.Client
	URI         string
}

type AppConfig struct {
	ServerConfig ServerConfig
	MongoConfig  MongoConfig
	JWT          struct {
		SecretKey string
	}
}
