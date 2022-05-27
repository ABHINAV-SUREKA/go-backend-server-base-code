# go-backend-server-base-code
### Description:
Base code for creating a Golang server/backend to quickly get started and build upon

### Items already included:
1. [**main.go**](https://github.com/ABHINAV-SUREKA/go-backend-server-base-code/blob/main/cmd/main.go#L9) - contains server configurations and starts the HTTP server
2. [**Gorilla Mux router**](https://github.com/ABHINAV-SUREKA/go-backend-server-base-code/blob/main/internal/app/routes.go#L11) - for handling routes
3. [**Route `'/status'`**](https://github.com/ABHINAV-SUREKA/go-backend-server-base-code/blob/main/internal/app/routes.go#L13) - to check status of the server
4. [**Middlewares:**](https://github.com/ABHINAV-SUREKA/go-backend-server-base-code/blob/main/internal/app/middleware.go)
   1. For all routes - enableCORS, logRequestWithDetails
   2. For desired routes - validateJWT
5. [**Utilities:**](https://github.com/ABHINAV-SUREKA/go-backend-server-base-code/blob/main/internal/app/utilities.go)
   1. writeJSON - write response to browser
   2. errorJSON - write error response to browser
   3. wrapMiddlewares - wrap a route with one or more middleware functions

### Steps to run the server:
1. Clone this repository
2. Run: `go run cmd/main.go`

_PS: Note the TODOs_
