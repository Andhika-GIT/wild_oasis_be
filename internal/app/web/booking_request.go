package web

import "time"

type EditBooking struct {
	NumGuests    *int   `json:"num_guests"`
	Observations string `json:"observations"`
}

type CreateBookingRequest struct {
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
}

type CreateBooking struct {
	StartDate    time.Time
	EndDate      time.Time
	NumNights    int
	NumGuests    int
	CabinPrice   float32
	ExtrasPrice  float32
	TotalPrice   float32
	Status       string
	HasBreakfast bool
	IsPaid       bool
	Observations string
	UserID       int
}
