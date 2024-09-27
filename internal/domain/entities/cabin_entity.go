package entities

import "time"

type Cabin struct {
	ID           int       `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt    time.Time `json:"-" gorm:"column:created_at;autoCreateTime"`
	Name         string    `json:"name" gorm:"column:name"`
	MaxCapacity  int       `json:"max_capacity" gorm:"column:max_capacity"`
	RegulerPrice int       `json:"regular_price" gorm:"column:regular_price"`
	Discount     int       `json:"discount,omitempty" gorm:"column:discount"`
	Description  string    `json:"description" gorm:"column:description"`
	Image        string    `json:"image" gorm:"column:image"`
}

func (a *Cabin) TableName() string {
	return "cabins"
}
