package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func ReadBodyRequest(request *http.Request, requestData any) error {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(requestData)

	if err != nil {
		return err
	}
	return nil
}

func SendResponse(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func SetCookie(w http.ResponseWriter, name string, value string, isProd bool) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

func ClearCookie(w http.ResponseWriter, name string, isProd bool) {
	cookie := http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return "", fmt.Errorf("cookie not found")
		default:
			return "", fmt.Errorf("server error")
		}
	}

	return cookie.Value, nil
}

func GetUserIDFromToken(r *http.Request) (int, error) {
	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		return 0, err
	}

	userIDRaw := claims["user_id"]

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return 0, err
	}

	return int(userIDFloat), nil
}
