package common

import (
	"net/http"
)

func WriteResponse(w http.ResponseWriter, response string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
	return
}