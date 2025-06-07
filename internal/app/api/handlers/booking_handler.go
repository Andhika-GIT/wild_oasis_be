package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/pkg/apperror"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type BookingHandler struct {
	BookingService *services.BookingService
	cabinService   *services.CabinService
}

func NewBookingHandler(bookingService *services.BookingService, cabinService *services.CabinService) *BookingHandler {
	return &BookingHandler{
		BookingService: bookingService,
		cabinService:   cabinService,
	}
}

func (c *BookingHandler) GetAllUserBookings(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromToken(r)

	if err != nil {
		utils.SendResponse(w, http.StatusUnauthorized, web.Response{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: "Unathorized",
		})
		return
	}

	bookings, err := c.BookingService.GetAllCurrentUserBookings(r.Context(), userID)

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

func (c *BookingHandler) GetSpesificUserBooking(w http.ResponseWriter, r *http.Request) {
	booking := r.Context().Value("booking").(entities.Booking)

	bookingData := web.ToBookingResponse(booking)

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully find booking",
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
			Message: fmt.Sprintf("Something went wrong, %s", err.Error()),
		})
		return
	}

	_, err = c.cabinService.FindById(r.Context(), id)
	if err != nil {
		utils.SendResponse(w, http.StatusNotFound, web.Response{
			Success: false,
			Code:    http.StatusNotFound,
			Message: "Cabin not found",
		})
		return
	}

	bookingResponse, err := c.BookingService.GetBookedDatesByCabinId(r.Context(), id)
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

func (c *BookingHandler) CreateUserBooking(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.CreateBookingRequest{}
	validate := validator.New()

	err := utils.ReadBodyRequest(r, bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong, %s", err.Error()),
		})
		return

	}

	err = validate.Struct(bodyRequest)

	if err != nil {
		errorResponse := apperror.ExtractValidationError(err)
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: strings.Join(errorResponse, ", "),
		})
		return
	}

	userID, err := utils.GetUserIDFromToken(r)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong, %s", err.Error()),
		})
		return
	}

	err = c.BookingService.CreateNewUserBooking(r.Context(), userID, bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong, %s", err.Error()),
		})
		return
	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully create reservation",
	})

}

func (c *BookingHandler) UpdateUserBooking(w http.ResponseWriter, r *http.Request) {
	booking := r.Context().Value("booking").(entities.Booking)

	bodyRequest := &web.EditBooking{}

	err := utils.ReadBodyRequest(r, bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong, %s", err.Error()),
		})
		return

	}

	if bodyRequest.Observations == "" || bodyRequest.NumGuests == nil {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "All fields are required",
		})
		return
	}

	err = c.BookingService.UpdateCurrentUserReservation(r.Context(), &booking, bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong, %s", err.Error()),
		})
		return

	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully update reservation",
	})
}

func (c *BookingHandler) DeleteCurrentUserBooking(w http.ResponseWriter, r *http.Request) {
	booking := r.Context().Value("booking").(entities.Booking)

	err := c.BookingService.DeleteCurrentUserBooking(r.Context(), booking)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong, %s", err.Error()),
		})
		return

	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Sucessfully delete booking",
	})
}
