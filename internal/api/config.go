package api

import "github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/db"

// Config for api/graphql operations
type Config interface { // we are exposing api.Config interface & its any required function outside the current package
}

type config struct {
	dbConfig *db.Config // creating dbConfig here coz api config receiver functions call db config receiver functions
}

func New(dbConfig *db.Config) Config {
	return &config{
		dbConfig: dbConfig,
	}
}
