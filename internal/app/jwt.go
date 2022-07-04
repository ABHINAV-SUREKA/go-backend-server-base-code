package app

import "flag"

// NewJWTSecretKey creates a new secret key for signing JWT
func NewJWTSecretKey() *string {
	jwtSecretKey := flag.String("jwt-secret-key", "some random secret key", "Secret key for signing JWT") // TODO: update jwt-secret-key value here
	return jwtSecretKey
}
