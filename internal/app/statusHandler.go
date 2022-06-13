package app

import (
	"github.com/ABHINAV-SUREKA/go-backend-server-base-code/pkg"
	"net/http"
)

type jsonResponse struct {
	Message string `json:"message"`
}

// statusHandler a handler function f(w,r) to check the status of the HTTP server
func (appConfig *config) statusHandler(w http.ResponseWriter, _ *http.Request) {
	statusResp := jsonResponse{
		Message: "Available",
	}

	if err := pkg.WriteJSON(w, http.StatusOK, statusResp, "status"); err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
}
