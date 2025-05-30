package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/go-chi/chi/v5"
)

func UserBookingMiddleware(bookingService *services.BookingService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userID, _ := utils.GetUserIDFromToken(r)

			paramsID := chi.URLParam(r, "bookingId")

			if paramsID == "" {
				utils.SendResponse(w, http.StatusNotFound, web.Response{
					Success: false,
					Code:    http.StatusNotFound,
					Message: "Cabin ID is required",
				})
				return
			}

			bookingID, err := strconv.Atoi(paramsID)

			if err != nil {
				utils.SendResponse(w, http.StatusInternalServerError, web.Response{
					Success: false,
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("something went wrong, %s", err.Error()),
				})
				return

			}

			_, err = bookingService.GetBookingById(r.Context(), bookingID)

			if err != nil {
				utils.SendResponse(w, http.StatusNotFound, web.Response{
					Success: false,
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
				return
			}

			booking, err := bookingService.CheckCurrentUserBooking(r.Context(), bookingID, userID)

			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized, web.Response{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "you are not allowed to run this action",
				})
				return
			}

			ctx := context.WithValue(r.Context(), "booking", booking)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
