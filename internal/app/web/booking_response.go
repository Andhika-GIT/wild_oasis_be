package web

import (
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
)

type BookingResponse struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	StartDate    time.Time `gorm:"column:start_date"`
	EndDate      time.Time `gorm:"column:end_date"`
	NumNights    int       `gorm:"column:num_nights"`
	NumGuests    int       `gorm:"column:num_guests"`
	CabinPrice   float32   `gorm:"column:cabin_price"`
	ExtrasPrice  float32   `gorm:"column:extras_price"`
	TotalPrice   float32   `gorm:"column:total_price"`
	Status       string    `gorm:"column:status"`
	HasBreakfast bool      `gorm:"column:has_breakfast"`
	IsPaid       bool      `gorm:"column:is_paid"`
	Observations string    `gorm:"column:observations"`
	CabinID      int       `gorm:"column:cabin_id"`
	UserID       int       `gorm:"column:user_id"`
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
		UserID:       booking.UserID,
	}
}

func ToBookingResponses(bookings []entities.Booking) []BookingResponse {
	var responses []BookingResponse

	for _, booking := range bookings {
		responses = append(responses, ToBookingResponse(booking))
	}

	return responses
}
