package main

import (
	"flag"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/constants"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/app"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopChan := make(chan os.Signal, 1) // create a channel to receive interrupts & shut down the server gracefully
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srvCfg := app.ServerConfig{}
	appCfg := app.AppConfig{}
	flag.IntVar(&srvCfg.Port, "port", 4000, "Server port to listen on")
	flag.DurationVar(&srvCfg.IdleTimeout, "idle-timeout", constants.IdleTimeout, "Maximum no. of seconds to wait for the next request when keep-alive is enabled")
	flag.DurationVar(&srvCfg.ReadTimeout, "read-timeout", constants.ReadTimeout, "Maximum no. of seconds before timing out reading of entire request, including the body")
	flag.DurationVar(&srvCfg.WriteTimeout, "write-timeout", constants.WriteTimeout, "Maximum no. of seconds before timing out writing of the response")
	flag.StringVar(&appCfg.JWT.SecretKey, "jwt-secret-key", "", "Secret key for signing JWT") // TODO: provide a secret key (say, a HMAC encrypted one) via cmd line arg
	flag.StringVar(&appCfg.MongoConfig.URI, "mongo-uri", "", "URI for establishing connection to mongoDB")
	flag.Parse()
	appCfg.ServerConfig = srvCfg

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})

	// Start http.Server
	go appCfg.RunHTTPServer(stopChan)

	<-stopChan
	log.Info("Server shut down gracefully")
}
