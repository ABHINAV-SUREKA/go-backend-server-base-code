package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Run starts HTTP server, listens to, and services the incoming user requests
func (appCfg *AppConfig) Run() {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", appCfg.ServerConfig.Port),
		Handler:      appCfg.Routes(),
		IdleTimeout:  appCfg.ServerConfig.IdleTimeout * time.Second,
		ReadTimeout:  appCfg.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: appCfg.ServerConfig.WriteTimeout * time.Second,
	}

	log.Infof("Starting server on port %v...", appCfg.ServerConfig.Port)
	log.Fatal(srv.ListenAndServe())
}
