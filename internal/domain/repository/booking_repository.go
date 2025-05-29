package repository

import (
	"context"
	"time"

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

func (r *BookingRepository) FindBookedDatesByCabinId(c context.Context, bookingDate time.Time, tx *gorm.DB, cabinId int, bookings *[]entities.Booking) error {
	err := tx.Where("cabin_id = ? AND (start_date >= ? OR status = ?)", cabinId, bookingDate, "checked-in").
		Find(&bookings).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) FindByUserId(c context.Context, tx *gorm.DB, userID int, bookings *[]entities.Booking) error {
	err := tx.Preload("Cabin", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "image")
	}).Where("user_id = ?", userID).Find(&bookings).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) FindById(c context.Context, tx *gorm.DB, bookingID int, booking *entities.Booking) error {
	err := tx.Where("id = ?", bookingID).First(&booking).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) Delete(c context.Context, tx *gorm.DB, booking *entities.Booking) error {
	err := tx.Delete(&booking).Error

	if err != nil {
		return err
	}

	return nil
}
