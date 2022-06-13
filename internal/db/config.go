package db

import (
	"context"
)

// Config interface declares methods for db operations
type Config interface { // we are exposing db.Config interface & its any required function outside the current package
	DBConnect(context.Context) error
}

// config struct implements Config interface methods
type config struct {
	dbClient interface{} // interface{} type enables config to work with client of any type of database
}

// New creates new db config struct
func New(dbClient interface{}) Config {
	return &config{
		dbClient: dbClient,
	}
}
