package main

import (
	"flag"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/api"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/app"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
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

	/* Create client for the database being used
	 */

	/* Configure HTTP server
	 */
	serverConfig := app.NewServerConfig()

	/* Create secret key for signing JWT
	 */
	jwtSecretKey := app.NewJWTSecretKey()

	/* Create configs in order of their dependencies
	 */
	dbConfig := db.New(&mongo.Client{})                         // for db operations // TODO: update the db client argument here
	apiConfig := api.New(dbConfig)                              // for api/graphql operations
	appConfig := app.New(serverConfig, apiConfig, jwtSecretKey) // for server, handler, etc. operations

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
