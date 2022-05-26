package main

import (
	"flag"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/constants"
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/internal/app"
)

func main() {
	srvCfg := app.ServerConfig{}
	appCfg := app.AppConfig{}
	flag.IntVar(&srvCfg.Port, "port", 4000, "Server port to listen on")
	flag.DurationVar(&srvCfg.IdleTimeout, "idle-timeout", constants.IdleTimeout, "Maximum no. of seconds to wait for the next request when keep-alive is enabled")
	flag.DurationVar(&srvCfg.ReadTimeout, "read-timeout", constants.ReadTimeout, "Maximum no. of seconds before timing out reading of entire request, including the body")
	flag.DurationVar(&srvCfg.WriteTimeout, "write-timeout", constants.WriteTimeout, "Maximum no. of seconds before timing out writing of the response")
	flag.StringVar(&appCfg.JWT.SecretKey, "jwt-secret-key", "", "JWT secret key for signing token") // TODO: provide a secret key (say, a HMAC encrypted one) via cmd line arg
	flag.Parse()
	appCfg.ServerConfig = srvCfg

	// appCfg.Run() starts http.Server
	appCfg.Run()
}
