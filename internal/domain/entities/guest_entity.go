package entities

import "time"

type Guest struct {
	ID          int64     `json:"id" gorm:"primary_key;column:id"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	FullName    string    `json:"full_name" gorm:"column:full_name"`
	Email       string    `json:"email" gorm:"column:email"`
	NationalID  string    `json:"national_id" gorm:"column:national_id"`
	Nationality string    `json:"nationality" gorm:"column:nationality"`
	CountryFlag string    `json:"country_flag" gorm:"column:country_flag"`
}

func (g *Guest) TableName() string {
	return "guests"
}
