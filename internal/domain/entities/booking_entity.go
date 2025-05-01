package entities

import "time"

type Booking struct {
	ID           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	StartDate    time.Time `json:"start_date" gorm:"column:start_date"`
	EndDate      time.Time `json:"end_date" gorm:"column:end_date"`
	NumNights    int       `json:"num_nights" gorm:"column:num_nights"`
	NumGuests    int       `json:"num_guests" gorm:"column:num_guests"`
	CabinPrice   float32   `json:"cabin_price" gorm:"column:cabin_price"`
	ExtrasPrice  float32   `json:"extras_price" gorm:"column:extras_price"`
	TotalPrice   float32   `json:"total_price" gorm:"column:total_price"`
	Status       string    `json:"status" gorm:"column:status"`
	HasBreakfast bool      `json:"has_breakfast" gorm:"column:has_breakfast"`
	IsPaid       bool      `json:"is_paid" gorm:"column:is_paid"`
	Observations string    `json:"observations" gorm:"column:observations"`
	CabinID      int       `json:"cabin_id" gorm:"column:cabin_id"`
	UserID       int       `json:"user_id" gorm:"column:user_id"`

	// Cabin Relation
	Cabin Cabin `gorm:"foreignKey:CabinID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// User Relation
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}

func (a *Booking) TableName() string {
	return "bookings"
}
