package entities

import "time"

type Cabin struct {
	ID           string    `gorm:"column:id;primaryKey"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	Name         string    `gorm:"column:name"`
	MaxCapacity  int       `gorm:"column:max_capacity"`
	RegulerPrice int       `gorm:"column:reguler_price"`
	Discount     int       `gorm:"column:discount"`
	Description  int       `gorm:"column:description"`
	Image        int       `gorm:"column:image"`
}
