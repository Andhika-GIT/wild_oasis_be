package entities

import "time"

type Cabin struct {
	ID           string    `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	Name         string    `json:"name" gorm:"column:name"`
	MaxCapacity  int       `json:"max_capacity" gorm:"column:max_capacity"`
	RegulerPrice int       `json:"regular_price" gorm:"column:reguler_price"`
	Discount     int       `json:"discount,omitempty" gorm:"column:discount"`
	Description  string    `json:"description" gorm:"column:description"`
	Image        string    `json:"image" gorm:"column:image"`
}

func (a *Cabin) TableName() string {
	return "cabins"
}
