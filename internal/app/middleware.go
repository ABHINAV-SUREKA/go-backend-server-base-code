package app

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Middleware
// a logic to be implemented on a user request before actually sending that request to the destined route (i.e., to the handler attached to that route)
// a function with a signature similar to f(http.Handler) http.Handler, which takes at least one handler and returns at least one new handler

// enableCORS tells browsers to enable go server running on different port than react server to accept requests from the latter
func (appCfg *AppConfig) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// enable CORS for the specified origin
		// TODO: 'Access-Control-Allow-Origin' value needs to be changed to something specific (such as your website URL) before moving to production
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// enable CORS for the specified headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		// return a success for 'OPTIONS' request type - imperative for Authorization Header to work
		if (*r).Method == "OPTIONS" {
			return
		}

		// implement 'next' handler
		next.ServeHTTP(w, r)
	})
}

// logRequestWithDetails logs each HTTP request with provided details
func (appCfg *AppConfig) logRequestWithDetails(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// log each request with provided 'Fields' before the request is handled
		log.WithFields(
			log.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"remoteAddr": r.RemoteAddr,
			}).Info("Handling request")

		// implement 'next' handler
		next.ServeHTTP(w, r)

		// log each request with provided 'Fields' after the request is handled
		log.WithFields(
			log.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"remoteAddr": r.RemoteAddr,
				"timeTaken":  time.Since(start),
			}).Info("Handled request")
	})
}

// checkJWT verifies JWT token sent in the request header by user
func (appCfg *AppConfig) validateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* Check for header authenticity
		 */
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			appCfg.errorJSON(w, errors.New("unauthorized - empty auth header"), http.StatusUnauthorized)
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			appCfg.errorJSON(w, errors.New("unauthorized - invalid auth header"), http.StatusUnauthorized)
			return
		}

		/* Check for jwt authenticity
		 */
		registeredClaims := jwt.StandardClaims{} // this will receive in all the registered/standard claims from JWT in request Auth Header
		// jwt.ParseWithClaims() takes Auth Bearer JWT from the request Header and a function f(Token) containing custom validation logic
		// it then does its own validation internally to ascertain token's format & integrity (including that of the received claims), as well as performs the validation via f(Token)
		// on success, it returns jwt.Token that contains: Signature (returned by f(Token)), Header, Claims, etc.
		token, err := jwt.ParseWithClaims(authHeaderParts[1], &registeredClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check for the expected HMAC algorithm
				return nil, fmt.Errorf("unauthorized - unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(appCfg.JWT.SecretKey), nil // return secret key (signature)
		})

		if token.Valid {
			// additional safety measures // maybe redundant
			if regexp.MustCompile(`^[0-9]+$`).MatchString(registeredClaims.Subject) {
				appCfg.errorJSON(w, errors.New("forbidden - invalid subject"), http.StatusForbidden)
				return
			}
			// TODO: 'Issuer' and 'Audience' values need to be updated as desired
			if registeredClaims.Issuer != "some issuer" {
				appCfg.errorJSON(w, errors.New("forbidden - invalid issuer"), http.StatusForbidden)
				return
			}
			if registeredClaims.Audience != "some audience" {
				appCfg.errorJSON(w, errors.New("forbidden - invalid audience"), http.StatusForbidden)
				return
			}

		} else if jwtErr, ok := err.(*jwt.ValidationError); ok {
			switch jwtErr.Errors {
			case jwt.ValidationErrorExpired, jwt.ValidationErrorNotValidYet:
				appCfg.errorJSON(w, err, http.StatusForbidden)
				break
			case jwt.ValidationErrorMalformed:
			case jwt.ValidationErrorClaimsInvalid:
			case jwt.ValidationErrorAudience:
			case jwt.ValidationErrorIssuer:
			case jwt.ValidationErrorIssuedAt:
			case jwt.ValidationErrorUnverifiable:
			default:
				appCfg.errorJSON(w, err, http.StatusUnauthorized)
				break
			}
			return
		} else {
			appCfg.errorJSON(w, fmt.Errorf("unauthorized - %s", err), http.StatusUnauthorized)
			return
		}

		// implement 'next' handler
		next.ServeHTTP(w, r)
	})
}
