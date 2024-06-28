package httpsvr

import (
	"encoding/json"
	"fmt"
	"log"
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

	if len(hostData.Records) == 0 {
		responseWithError(w, 400, fmt.Sprintf("Operation Failed: Host records can not be empty"))
		return
	}

	env := chi.URLParam(r, "env")
	log.Printf("Environment: %s\n", env)

	hostManager, err := h.hostFactory.GetManager(env)
	log.Printf("Host Manager: %v\n", hostManager)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("%v", err))
		return
	}

	response, err := hostManager.AddSubDomain(&hostData)
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

	if len(hostData.Records) == 0 {
		responseWithError(w, 400, fmt.Sprintf("Operation Failed: Host records can not be empty"))
		return
	}

	env := chi.URLParam(r, "env")
	hostManager, err := h.hostFactory.GetManager(env)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("%v", err))
		return
	}

	response, err := hostManager.DeleteSubDomain(&hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error Updating Sub-domain: %v", err))
		return
	}

	responseWithJSON(w, 200, response)
}

func (h HttpAPIServer) listSubDomain(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	hostData := internals.HostData{}

	err := decoder.Decode(&hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	env := chi.URLParam(r, "env")
	hostManager, err := h.hostFactory.GetManager(env)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("%v", err))
		return
	}

	response, err := hostManager.ListSubDomain(&hostData)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error Updating Sub-domain: %v", err))
		return
	}

	responseWithJSON(w, 200, response)
}
