package repository

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"gorm.io/gorm"
)

type BookingRepository struct{}

func (r *BookingRepository) Create(c context.Context, tx *gorm.DB, booking *entities.Booking) error {
	err := tx.Create(&booking).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) FindAll(c context.Context, tx *gorm.DB, bookings *[]entities.Booking) error {
	err := tx.Find(&bookings).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) FindBookedDatesByCabinId(c context.Context, bookingDate string, tx *gorm.DB, cabinId int, bookings *[]entities.Booking) error {
	err := tx.Where("id = ?", cabinId).Where("start_date >= ?", bookingDate).Or("status = ?", "checked-in").Find(&bookings).Error

	if err != nil {
		return err
	}

	return nil
}
