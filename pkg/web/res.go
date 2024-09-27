package web

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
