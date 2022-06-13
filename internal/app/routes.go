package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Routes multiplexes/routes requests to the appropriate route
func (appConfig *config) Routes() http.Handler {
	// mux.NewRouter() returns *Router, which is of type http.Handler interface since *Router implements ServeHTTP(w,r) method
	handler := mux.NewRouter()
	// handler.HandleFunc() converts the passed in func(w,r) to http.HandlerFunc (a http.Handler interface) and calls http.Handle on it
	handler.HandleFunc("/status", appConfig.statusHandler).Methods("GET")

	/* Add more routes here (wrap them in middleware functions, if needed, using wrapMiddlewares)
	 */

	// handler.Use() wraps all routes with the specified middlewares
	// mux.CORSMethodMiddleware() sets the Access-Control-Allow-Methods header in the response for a route's allowed methods/request types, iff there is 'OPTIONS' method in the allowed methods
	handler.Use(mux.CORSMethodMiddleware(handler), appConfig.enableCORS, appConfig.logRequestWithDetails)
	return handler
}
