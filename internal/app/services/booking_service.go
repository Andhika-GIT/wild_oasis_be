package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/Andhika-GIT/wild_oasis_be/pkg/file"
	"gorm.io/gorm"
)

type BookingService struct {
	repository *repository.BookingRepository
	DB         *gorm.DB
}

func NewBookingService(repository *repository.BookingRepository, DB *gorm.DB) *BookingService {
	return &BookingService{
		repository: repository,
		DB:         DB,
	}
}

func (s *BookingService) GetBookedDatesByCabinId(c context.Context, cabinId int) ([]web.BookingResponse, error) {
	var bookings []entities.Booking

	tx := s.DB.WithContext(c).Begin()
	defer tx.Rollback()

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	todayStr := today.Format(time.RFC3339) // ex Output: "2025-01-02T00:00:00Z"

	err := s.repository.FindBookedDatesByCabinId(c, todayStr, tx, cabinId, &bookings)

	if err != nil {
		return []web.BookingResponse{}, fmt.Errorf("error while find bookings : %v", err)
	}

	bookingResponses := web.ToBookingResponses(bookings)

	return bookingResponses, tx.Commit().Error
}

func (s *BookingService) SeedBookings(c context.Context) error {

	tx := s.DB.WithContext(c).Begin()

	// rollback after all function done
	defer tx.Rollback()

	// read file from json
	bookings, err := file.LoadFromJsonFile[[]entities.Booking]("./data/bookings.json")
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	fmt.Println("Bookings loaded from JSON:")
	for _, booking := range bookings {
		fmt.Printf("Booking: %+v\n", booking)
	}

	//  reset all data first first
	err = tx.Exec("DELETE from bookings").Error
	if err != nil {
		return fmt.Errorf("error when deleting all bookings : %v", err)
	}

	err = tx.Exec("TRUNCATE TABLE bookings RESTART IDENTITY CASCADE").Error
	if err != nil {
		return fmt.Errorf("error when truncating table: %v", err)
	}

	for _, booking := range bookings {
		err = s.repository.Create(c, tx, &booking)

		if err != nil {
			return fmt.Errorf("error create booking : %v", err)
		}
	}

	return tx.Commit().Error

}
