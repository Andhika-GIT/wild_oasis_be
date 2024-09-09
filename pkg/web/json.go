package web

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
