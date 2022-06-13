package api

// Config interface declares functions for api/graphql operations
type Config interface { // we are exposing api.Config interface & its any required function outside the current package
}

// config struct implements Config interface methods
type config struct {
	dbConfig interface{} // creating dbConfig here coz api config receiver functions call db config receiver functions
}

// New creates new api config struct
func New(dbConfig interface{}) Config {
	return &config{
		dbConfig: dbConfig,
	}
}
