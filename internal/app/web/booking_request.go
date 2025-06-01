package web

type EditReservation struct {
	NumGuests    *int   `json:"num_guests"`
	Observations string `json:"observations"`
}
