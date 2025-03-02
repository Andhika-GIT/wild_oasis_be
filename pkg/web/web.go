package web

import (
	"encoding/json"
	"net/http"
	"time"
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

func SetCookie(w http.ResponseWriter, token string, isProd bool) {
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}
