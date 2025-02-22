package entities

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	Email     string    `json:"email" gorm:"column:email;not null"`
	Password  string    `gorm:"column:password; not null"`
	GuestID   *int      `json:"guest_id" gorm:"column:guest_id"`

	// Guest Relation
	Guest Guest `gorm:"foreignKey:GuestID;references:ID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}

func (a *User) TableName() string {
	return "users"
}
