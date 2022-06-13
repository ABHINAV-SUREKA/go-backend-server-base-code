package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// RunHTTPServer starts HTTP server, listens to, and services the incoming user requests
func (appConfig *config) RunHTTPServer() error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", appConfig.serverConfig.port),
		Handler:      appConfig.Routes(),
		IdleTimeout:  appConfig.serverConfig.idleTimeout * time.Second,
		ReadTimeout:  appConfig.serverConfig.readTimeout * time.Second,
		WriteTimeout: appConfig.serverConfig.writeTimeout * time.Second,
	}

	log.Infof("Running HTTP server on port %v...", appConfig.serverConfig.port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Errorf("error starting server: %s", err.Error())
		return err
	}
	return nil
}
