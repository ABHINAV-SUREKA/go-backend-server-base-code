package app

import (
	"flag"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/constants"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/api"
	"time"
)

// Config interface declaring functions for server, handler, etc. operations
type Config interface { // we are exposing app.Config interface & its any required function outside the current package
	RunHTTPServer() error
}

type serverConfig struct {
	port         int
	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// config struct implements Config interface methods
type config struct {
	serverConfig *serverConfig
	apiConfig    *api.Config // creating apiConfig here coz app config receiver functions call api config receiver functions
	secretKey    string
}

// New creates a new app config struct
func New(apiConfig *api.Config) Config {
	port := flag.Int("port", 4000, "Server port to listen on") // can also user IntVar in main() as well
	idleTimeout := flag.Duration("idle-timeout", constants.IdleTimeout, "Maximum no. of seconds to wait for the next request when keep-alive is enabled")
	readTimeout := flag.Duration("read-timeout", constants.ReadTimeout, "Maximum no. of seconds before timing out reading of entire request, including the body")
	writeTimeout := flag.Duration("write-timeout", constants.WriteTimeout, "Maximum no. of seconds before timing out writing of the response")
	jwtSecretKey := flag.String("jwt-secret-key", "some secret key", "Secret key for signing JWT") // TODO: update secret key (say, a HMAC encrypted one) via cmd line arg

	return &config{
		serverConfig: &serverConfig{
			port:         *port,
			idleTimeout:  *idleTimeout,
			readTimeout:  *readTimeout,
			writeTimeout: *writeTimeout,
		},
		secretKey: *jwtSecretKey,
		apiConfig: apiConfig,
	}
}
