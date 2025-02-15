package entities

import "time"

type Settings struct {
	ID                int64     `json:"id" gorm:"primary_key;column:id"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	MinBookingLength  int       `json:"min_booking_length" gorm:"column:min_booking_length"`
	MaxBookingLength  int       `json:"max_booking_length" gorm:"column:max_booking_length"`
	MaxGuestsPerCabin int       `json:"max_guests_per_cabin" gorm:"column:max_guests_per_cabin"`
	BreakfastPrice    float32   `json:"breakfast_price" gorm:"column:breakfast_price"`
}

func (s *Settings) TableName() string {
	return "settings"
}
