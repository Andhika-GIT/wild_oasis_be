package web

type UserResponse struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	GuestID *int   `json:"guest_id"`
}
