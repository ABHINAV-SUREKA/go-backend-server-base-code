package main

import (
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/api"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/app"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// initConfigs creates/initialises all the structs and interfaces throughout our program in order of their intrinsic
// this function can also be refactored to use dependency injection library such as 'github.com/google/wire' in it, if the existing dependency graph becomes complex
func initConfigs() (app.Config, db.Config) {
	/* Create client here for the database being used
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

	return appConfig, dbConfig
}
