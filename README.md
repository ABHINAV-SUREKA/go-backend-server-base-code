# go-backend-server-base-code
Base code for a Golang server to quickly get started and build upon

### Items already included:
1. **Gorilla Mux router** - for handling routes
2. **Route `'/status'`** - to check status of the server
3. **Middlewares:**
   1. For all routes - enableCORS
   2. For desired routes - validateJWT
4. **Utilities:**
   1. writeJSON - write response to browser
   2. errorJSON - write error response to browser
   3. wrapMiddlewares - wrap a route with one or more middleware functions

### Steps to run the server:
1. Clone this repository
2. Run: `go run cmd/main.go`

_PS: Note the TODOs_
