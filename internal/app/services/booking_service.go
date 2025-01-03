package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
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
