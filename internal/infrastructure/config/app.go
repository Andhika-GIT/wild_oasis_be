package config

import (
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/api/handlers"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Router            *chi.Mux
	CabinService      *services.CabinService
	BookingService    *services.BookingService
	SettingService    *services.SettingService
	CloudinaryService *services.CloudinaryService
	AuthService       *services.AuthService
	Global            *viper.Viper
}

func Bootstrap() *AppConfig {
	v := NewViper()
	l := NewLogger()
	db := NewDatabase(v, &l)

	// repository
	cabinRepository := repository.CabinRepository{}
	bookingRepository := repository.BookingRepository{}
	settingRepository := repository.SettingRepository{}
	userRepository := repository.UserRepository{}

	// services
	cloudinaryService := services.NewCloudinaryService(v)
	cabinService := services.NewCabinService(&cabinRepository, db, cloudinaryService)
	settingService := services.NewSettingService(&settingRepository, db)
	bookingService := services.NewBookingService(&bookingRepository, db)
	authService := services.NewAuthService(&userRepository, db, v)

	// handlers
	cabinHandler := handlers.NewCabinHandler(cabinService)
	bookingHandler := handlers.NewBookingHandler(bookingService, cabinService)
	settingHandler := handlers.NewSettingHandler(settingService)
	cloudinaryHandler := handlers.NewCloudinaryHandler(cloudinaryService)
	authHandler := handlers.NewAuthHandler(authService, v)

	router := handlers.NewRouter(cabinHandler, bookingHandler, settingHandler, cloudinaryHandler, authHandler, v)

	return &AppConfig{
		Router:            router.GetRouter(),
		CabinService:      cabinService,
		BookingService:    bookingService,
		SettingService:    settingService,
		CloudinaryService: cloudinaryService,
		AuthService:       authService,
		Global:            v,
	}
}
