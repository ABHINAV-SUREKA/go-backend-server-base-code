package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type jsonError struct {
	Message string `json:"message"`
}

// writeJSON writes response content to browser
func (appCfg *AppConfig) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	dataWrapper := make(map[string]interface{})
	dataWrapper[wrap] = data

	byteArr, err := json.MarshalIndent(dataWrapper, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(byteArr)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON writes error response content to browser
func (appCfg *AppConfig) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	errorResp := jsonError{
		Message: err.Error(),
	}

	if err := appCfg.writeJSON(w, statusCode, errorResp, "error"); err != nil {
		log.Errorf("writeJSON() failed to write errorResp: %v | error: %s", errorResp.Message, err.Error())
	}
}

// wrapMiddlewares wraps one or more middleware functions around a handler for f(w,r)
// thus, middlewares can be selectively wrapped on individual routes
func (appCfg *AppConfig) wrapMiddlewares(function func(w http.ResponseWriter, r *http.Request), mwFuncs ...func(http.Handler) http.Handler) http.Handler {
	// http.HandlerFunc() converts f(w,r) to http.HandlerFunc (a http.Handler interface)
	handlerFunc := http.HandlerFunc(function)
	// convertHandlerFuncToHandler() converts http.HandlerFunc to http.Handler
	handler := convertHandlerFuncToHandler(handlerFunc)

	for i := range mwFuncs {
		handler = mwFuncs[len(mwFuncs)-1-i](handler) // The index is reversed so that the last middleware in the list of middlewares is the first to get wrapped around the function(w,r) i.e., (gets executed last)
	}
	return handler
}

// convertHandlerFuncToHandler converts http.HandlerFunc to http.Handler
func convertHandlerFuncToHandler(handlerFunc http.HandlerFunc) http.Handler {
	return handlerFunc
}
