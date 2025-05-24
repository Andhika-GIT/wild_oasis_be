package web

import (
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
)

type BookingResponse struct {
	ID           int64     `json:"id"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	NumNights    int       `json:"num_nights"`
	NumGuests    int       `json:"num_guests"`
	CabinPrice   float32   `json:"cabin_price"`
	ExtrasPrice  float32   `json:"extras_price"`
	TotalPrice   float32   `json:"total_price"`
	Status       string    `json:"status"`
	HasBreakfast bool      `json:"has_breakfast"`
	IsPaid       bool      `json:"is_paid"`
	Observations string    `json:"observations"`
	CabinID      int       `json:"cabin_id"`
}

func ToBookingResponse(booking entities.Booking) BookingResponse {
	return BookingResponse{
		ID:           booking.ID,
		StartDate:    booking.StartDate,
		EndDate:      booking.EndDate,
		NumNights:    booking.NumNights,
		NumGuests:    booking.NumGuests,
		CabinPrice:   booking.CabinPrice,
		ExtrasPrice:  booking.ExtrasPrice,
		TotalPrice:   booking.TotalPrice,
		Status:       booking.Status,
		HasBreakfast: booking.HasBreakfast,
		IsPaid:       booking.IsPaid,
		Observations: booking.Observations,
		CabinID:      booking.CabinID,
	}
}

func ToBookingResponses(bookings []entities.Booking) []BookingResponse {
	var responses []BookingResponse

	for _, booking := range bookings {
		responses = append(responses, ToBookingResponse(booking))
	}

	return responses
}
