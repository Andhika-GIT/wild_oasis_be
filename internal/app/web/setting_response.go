package web

import "github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"

type SettingResponse struct {
	ID                int64   `json:"id"`
	MinBookingLength  int     `json:"min_booking_length"`
	MaxBookingLength  int     `json:"max_booking_length"`
	MaxGuestsPerCabin int     `json:"max_guests_per_cabin"`
	BreakfastPrice    float32 `json:"breakfast_price"`
}

func ToSettingResponse(setting entities.Settings) SettingResponse {
	return SettingResponse{
		ID:               setting.ID,
		MinBookingLength: setting.MinBookingLength,
		MaxBookingLength: setting.MaxBookingLength,
		BreakfastPrice:   setting.BreakfastPrice,
	}
}
