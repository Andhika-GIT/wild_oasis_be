package handlers

import (
	"net/http"
	"strconv"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/go-chi/chi/v5"
)

type BookingHandler struct {
	bookingService *services.BookingService
	cabinService   *services.CabinService
}

func NewBookingHandler(bookingService *services.BookingService, cabinService *services.CabinService) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
		cabinService:   cabinService,
	}
}

func (c *BookingHandler) GetCurrentUserBooking(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromToken(r)

	if err != nil {
		utils.SendResponse(w, http.StatusUnauthorized, web.Response{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: "Unathorized",
		})
		return
	}

	bookings, err := c.bookingService.FindCurrentUserBooking(r.Context(), userID)

	if err != nil {
		utils.SendResponse(w, http.StatusNotFound, web.Response{
			Success: false,
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	bookingData := web.ToBookingResponses(bookings)

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Sucessfully found user booking",
		Data:    bookingData,
	})
}

func (c *BookingHandler) GetBookedDatesByCabinId(w http.ResponseWriter, r *http.Request) {
	cabinId := chi.URLParam(r, "cabinId")

	if cabinId == "" {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Cabin id is required",
		})
		return
	}

	id, err := strconv.Atoi(cabinId)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		return
	}

	_, err = c.cabinService.FindById(r.Context(), id)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound, web.Response{
			Success: false,
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	bookingResponse, err := c.bookingService.GetBookedDatesByCabinId(r.Context(), id)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound, web.Response{
			Success: false,
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully find booking",
		Data:    bookingResponse,
	})
}
