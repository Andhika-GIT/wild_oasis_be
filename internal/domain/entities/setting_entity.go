package entities

import "time"

type Settings struct {
	ID                int64     `gorm:"primary_key;column:id"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime"`
	MinBookingLength  int       `gorm:"column:min_booking_length"`
	MaxBookingLength  int       `gorm:"column:max_booking_length"`
	MaxGuestsPerCabin int       `gorm:"column:max_guests_per_cabin"`
	BreakfastPrice    float32   `gorm:"column:breakfast_price"`
}

func (s *Settings) TableName() string {
	return "settings"
}
