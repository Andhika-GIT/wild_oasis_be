package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/Andhika-GIT/wild_oasis_be/pkg/date"
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

func (s *BookingService) GetBookedDatesByCabinId(c context.Context, cabinId int) ([]string, error) {
	var bookings []entities.Booking
	var bookingDates [][]string
	var formattedDates []string

	tx := s.DB.WithContext(c).Begin()
	defer tx.Rollback()

	now := time.Now().UTC()                                                      // ex output : 2025-01-02T14:23:45Z
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // ex output : 2025-01-02 00:00:00 +0000 UTC
	// todayStr := today.Format(time.RFC3339) // ex Output: "2025-01-02T00:00:00Z"

	err := s.repository.FindBookedDatesByCabinId(c, today, tx, cabinId, &bookings)

	if err != nil {
		return []string{}, fmt.Errorf("error while find bookings : %v", err)
	}

	bookingResponses := web.ToBookingResponses(bookings)

	for _, booking := range bookingResponses {
		bookingDates = append(bookingDates, date.EachDayOfInterval(booking.StartDate, booking.EndDate))
	}

	fmt.Println("---- INI ADALAH BOOKING DATE ----")
	for _, bookingDate := range bookingDates {
		for _, date := range bookingDate {
			formattedDates = append(formattedDates, date)
		}
	}

	return formattedDates, nil

}

func (s *BookingService) FindCurrentUserBooking(c context.Context, userID int) ([]entities.Booking, error) {
	var bookings []entities.Booking

	tx := s.DB.WithContext(c)

	err := s.repository.FindByUserId(c, tx, userID, &bookings)

	if err != nil {
		return bookings, fmt.Errorf("error when finding user booking %v", err)
	}

	return bookings, nil

}

func (s *BookingService) DeleteCurrentUserBooking(c context.Context, userID int, bookingID int) (int, string, bool) {
	var booking entities.Booking

	tx := s.DB.WithContext(c)

	_, err := s.FindCurrentUserBooking(c, userID)

	if err != nil {
		return http.StatusUnauthorized, fmt.Sprint("you are not allowed to run this action"), false
	}

	err = s.repository.FindById(c, tx, bookingID, &booking)

	if err != nil {
		return http.StatusNotFound, fmt.Sprint("booking not found"), false
	}

	err = s.repository.Delete(c, tx, &booking)

	if err != nil {
		return http.StatusInternalServerError, err.Error(), false
	}

	return http.StatusOK, "Sucessfully delete user booking", true

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
