package config

import (
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/api/handlers"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/go-chi/chi/v5"
)

type AppConfig struct {
	Router            *chi.Mux
	CabinService      *services.CabinService
	BookingService    *services.BookingService
	CloudinaryService *services.CloudinaryService
}

func Bootstrap() *AppConfig {
	v := NewViper()
	l := NewLogger()
	db := NewDatabase(v, &l)

	// repository
	cabinRepository := repository.CabinRepository{}
	bookingRepository := repository.BookingRepository{}

	// services
	cloudinaryService := services.NewCloudinaryService(v)
	cabinService := services.NewCabinService(&cabinRepository, db, cloudinaryService)
	bookingService := services.NewBookingService(&bookingRepository, db)

	// handlers
	cabinHandler := handlers.NewCabinHandler(cabinService)
	bookingHandler := handlers.NewBookingHandler(bookingService, cabinService)
	cloudinaryHandler := handlers.NewCloudinaryHandler(cloudinaryService)

	router := handlers.NewRouter(cabinHandler, bookingHandler, cloudinaryHandler)

	return &AppConfig{
		Router:            router.GetRouter(),
		CabinService:      cabinService,
		BookingService:    bookingService,
		CloudinaryService: cloudinaryService,
	}
}
