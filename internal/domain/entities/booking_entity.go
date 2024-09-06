package entities

import "time"

type Booking struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
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
	GuestID      int       `gorm:"column:guest_id"`

	// Cabin Relation
	Cabin Cabin `gorm:"foreignKey:CabinID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Guest Relation
	Guest Guest `gorm:"foreignKey:GuestID;references:ID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}

func (a *Booking) TableName() string {
	return "bookings"
}
