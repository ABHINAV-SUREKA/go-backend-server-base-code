package app

import (
	"time"
)

type ServerConfig struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppConfig struct {
	ServerConfig ServerConfig
	JWT          struct {
		SecretKey string
	}
}
