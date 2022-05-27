package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run starts HTTP server, listens to, and services the incoming user requests
func (appCfg *AppConfig) Run() {
	stopChan := make(chan os.Signal, 1) // create a channel to receive interrupts & shut down the server gracefully
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", appCfg.ServerConfig.Port),
		Handler:      appCfg.Routes(),
		IdleTimeout:  appCfg.ServerConfig.IdleTimeout * time.Second,
		ReadTimeout:  appCfg.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: appCfg.ServerConfig.WriteTimeout * time.Second,
	}

	log.Infof("Starting server on port %v...", appCfg.ServerConfig.Port)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Errorf("error starting server: %s", err.Error())
			stopChan <- syscall.SIGTERM
		}
	}()

	<-stopChan
	log.Info("Server shut down gracefully")
}
