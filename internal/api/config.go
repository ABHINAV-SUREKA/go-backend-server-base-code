package api

import "github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/db"

// Config interface declaring functions for api/graphql operations
type Config interface { // we are exposing api.Config interface & its any required function outside the current package
}

// config struct implements Config interface methods
type config struct {
	dbConfig *db.Config // creating dbConfig here coz api config receiver functions call db config receiver functions
}

// New creates new api config struct
func New(dbConfig *db.Config) Config {
	return &config{
		dbConfig: dbConfig,
	}
}
