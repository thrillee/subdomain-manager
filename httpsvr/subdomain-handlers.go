package httpsvr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thrillee/namecheap-dns-manager/internals"
)

func (h HttpAPIServer) addSubDomain(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	hostData := internals.HostData{}

	err := decoder.Decode(&hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	env := chi.URLParam(r, "env")
	isLive := env == "prod"

	hostManager := h.getHostManagerService(isLive)
	response, err := hostManager.AddSubDomain(hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error Updating Sub-domain: %v", err))
		return
	}

	responseWithJSON(w, 200, response)
}

func (h HttpAPIServer) deleteSubDomain(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	hostData := internals.HostData{}

	err := decoder.Decode(&hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	env := chi.URLParam(r, "env")
	isLive := env == "prod"

	hostManager := h.getHostManagerService(isLive)
	response, err := hostManager.DeleteSubDomain(hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error Updating Sub-domain: %v", err))
		return
	}

	responseWithJSON(w, 200, response)
}
