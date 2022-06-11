package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"syscall"
	"time"
)

// RunHTTPServer starts HTTP server, listens to, and services the incoming user requests
func (appCfg *AppConfig) RunHTTPServer(stopChan chan os.Signal) {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", appCfg.ServerConfig.Port),
		Handler:      appCfg.Routes(),
		IdleTimeout:  appCfg.ServerConfig.IdleTimeout * time.Second,
		ReadTimeout:  appCfg.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: appCfg.ServerConfig.WriteTimeout * time.Second,
	}

	log.Infof("Running HTTP server on port %v...", appCfg.ServerConfig.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Errorf("error starting server: %s", err.Error())
		stopChan <- syscall.SIGTERM
	}
}
