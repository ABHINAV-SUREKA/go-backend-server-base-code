package app

import (
	"flag"
	"fmt"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/constants"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// NewServerConfig creates configuration the HTTP server will run with
func NewServerConfig() *serverConfig {
	port := flag.Int("port", 4000, "Server port to listen on") // can also user IntVar in main() as well
	idleTimeout := flag.Duration("idle-timeout", constants.IdleTimeout, "Maximum no. of seconds to wait for the next request when keep-alive is enabled")
	readTimeout := flag.Duration("read-timeout", constants.ReadTimeout, "Maximum no. of seconds before timing out reading of entire request, including the body")
	writeTimeout := flag.Duration("write-timeout", constants.WriteTimeout, "Maximum no. of seconds before timing out writing of the response")

	return &serverConfig{
		port:         *port,
		idleTimeout:  *idleTimeout,
		readTimeout:  *readTimeout,
		writeTimeout: *writeTimeout,
	}
}

// RunHTTPServer starts HTTP server, listens to, and services the incoming user requests
func (appConfig *config) RunHTTPServer() error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", appConfig.serverConfig.(*serverConfig).port),
		Handler:      appConfig.routes(),
		IdleTimeout:  appConfig.serverConfig.(*serverConfig).idleTimeout * time.Second,
		ReadTimeout:  appConfig.serverConfig.(*serverConfig).readTimeout * time.Second,
		WriteTimeout: appConfig.serverConfig.(*serverConfig).writeTimeout * time.Second,
	}

	log.Infof("Starting HTTP server on port %v...", appConfig.serverConfig.(*serverConfig).port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Errorf("error starting server: %s", err.Error())
		return err
	}
	return nil
}
