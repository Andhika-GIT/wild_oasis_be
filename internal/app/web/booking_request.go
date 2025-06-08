package web

import (
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
)

type EditBooking struct {
	NumGuests    *int   `json:"num_guests"`
	Observations string `json:"observations"`
}

type CreateBookingRequest struct {
	StartDate    time.Time `json:"start_date" validate:"required"`
	EndDate      time.Time `json:"end_date" validate:"required"`
	NumNights    int       `json:"num_nights" validate:"required"`
	NumGuests    int       `json:"num_guests" validate:"required"`
	CabinPrice   float32   `json:"cabin_price" validate:"gte=0"`
	ExtrasPrice  float32   `json:"extras_price" validate:"gte=0"`
	TotalPrice   float32   `json:"total_price" validate:"gte=0"`
	Status       string    `json:"status" validate:"required"`
	HasBreakfast bool      `json:"has_breakfast" ` // boolean biasanya optional
	IsPaid       bool      `json:"is_paid"`
	Observations string    `json:"observations" validate:"required"`
	CabinID      int       `json:"cabin_id" validate:"required"`
}

func ToBookingEntity(bookingData *CreateBookingRequest, userID int) *entities.Booking {
	return &entities.Booking{
		StartDate:    bookingData.StartDate,
		EndDate:      bookingData.EndDate,
		NumNights:    bookingData.NumNights,
		NumGuests:    bookingData.NumGuests,
		CabinPrice:   bookingData.CabinPrice,
		ExtrasPrice:  bookingData.ExtrasPrice,
		TotalPrice:   bookingData.TotalPrice,
		Status:       bookingData.Status,
		HasBreakfast: bookingData.HasBreakfast,
		IsPaid:       bookingData.IsPaid,
		Observations: bookingData.Observations,
		CabinID:      bookingData.CabinID,
		UserID:       userID,
	}
}
