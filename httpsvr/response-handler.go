package httpsvr

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Response with 5XX error: ", msg)
	}

	type errResponse struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	}

	responseWithJSON(w, code, errResponse{
		Error:      msg,
		StatusCode: code,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
