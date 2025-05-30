package web

import (
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
)

type BookedCabin struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	MaxCapacity int    `json:"max_capacity"`
}

type BookingResponse struct {
	ID           int64       `json:"id"`
	StartDate    time.Time   `json:"start_date"`
	EndDate      time.Time   `json:"end_date"`
	NumNights    int         `json:"num_nights"`
	NumGuests    int         `json:"num_guests"`
	CabinPrice   float32     `json:"cabin_price"`
	ExtrasPrice  float32     `json:"extras_price"`
	TotalPrice   float32     `json:"total_price"`
	Status       string      `json:"status"`
	HasBreakfast bool        `json:"has_breakfast"`
	IsPaid       bool        `json:"is_paid"`
	Observations string      `json:"observations"`
	CreatedAt    time.Time   `json:"created_at"`
	Cabin        BookedCabin `json:"cabin"`
}

func ToBookingResponse(booking entities.Booking) BookingResponse {
	Cabin := &BookedCabin{
		Name:        booking.Cabin.Name,
		Image:       booking.Cabin.Image,
		MaxCapacity: booking.Cabin.MaxCapacity,
	}

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
		CreatedAt:    booking.CreatedAt,
		Cabin:        *Cabin,
	}
}

func ToBookingResponses(bookings []entities.Booking) []BookingResponse {
	var responses []BookingResponse

	for _, booking := range bookings {
		responses = append(responses, ToBookingResponse(booking))
	}

	return responses
}
