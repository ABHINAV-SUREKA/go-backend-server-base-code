package app

import (
	"net/http"
)

type jsonResponse struct {
	Message string `json:"message"`
}

// statusHandler a handler function f(w,r) to check the status of the HTTP server
func (appCfg *AppConfig) statusHandler(w http.ResponseWriter, _ *http.Request) {
	statusResp := jsonResponse{
		Message: "Available",
	}

	if err := appCfg.writeJSON(w, http.StatusOK, statusResp, "status"); err != nil {
		appCfg.errorJSON(w, err)
		return
	}
}
