package web

import (
	"encoding/json"
	"net/http"
)

func ReadBodyRequest(request *http.Request, requestData interface{}) error {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(requestData)

	if err != nil {
		return err
	}
	return nil
}

func SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
