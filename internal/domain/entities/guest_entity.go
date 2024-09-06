package entities

import "time"

type Guest struct {
	ID          int64     `gorm:"primary_key;column:id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	FullName    string    `gorm:"column:full_name"`
	Email       string    `gorm:"column:email"`
	NationalID  string    `gorm:"column:national_id"`
	Nationality string    `gorm:"column:nationality"`
	CountryFlag string    `gorm:"column:country_flag"`
}

func (g *Guest) TableName() string {
	return "guests"
}
