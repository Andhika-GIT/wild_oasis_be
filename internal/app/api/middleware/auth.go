package middleware

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/go-chi/jwtauth/v5"
)

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)

			tokenStr, err := utils.GetCookie(w, r, "access_token")
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized, web.Response{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				})
				return
			}

			token, err := jwtauth.VerifyToken(tokenAuth, tokenStr)
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized, web.Response{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				})
				return
			}

			ctx := jwtauth.NewContext(r.Context(), token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
