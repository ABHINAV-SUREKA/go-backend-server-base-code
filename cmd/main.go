package main

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	/* Parse cmd line arguments
	 */
	flag.Parse()

	/* Create channel for interrupt signals
	 */
	stopChan := make(chan os.Signal, 1) // create a channel to receive interrupts & shut down the server gracefully
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	/* Set logging format
	 */
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})

	/* Create a new context
	 */
	_, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	/* Create/initialise all the configs (structs and interfaces) used throughout the program
	 */
	appConfig, _ := initConfigs()

	/* Close database connection here (using defer preferably)
	 */

	/* Connect to database server here
	 */

	/* Start HTTP server
	 */
	go func() {
		err := appConfig.RunHTTPServer()
		if err != nil {
			stopChan <- syscall.SIGTERM
		}
	}()

	log.Infof("Received: %v. Server shut down gracefully", <-stopChan)
}
