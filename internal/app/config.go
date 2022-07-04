package app

import (
	"net/http"
	"time"
)

// Config interface declares functions for server, handler, etc. operations
type Config interface { // we are exposing app.Config interface & its any required function outside the current package
	RunHTTPServer() error
	routes() http.Handler
	statusHandler(http.ResponseWriter, *http.Request)
}

type serverConfig struct {
	port         int
	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// config struct implements Config interface methods
type config struct {
	serverConfig interface{}
	apiConfig    interface{} // declaring apiConfig here coz app config receiver functions call api config receiver functions
	jwtSecretKey *string
}

// New creates a new app config struct
func New(serverConfig interface{}, apiConfig interface{}, jwtSecretKey *string) Config {
	return &config{
		serverConfig: serverConfig,
		apiConfig:    apiConfig,
		jwtSecretKey: jwtSecretKey,
	}
}
