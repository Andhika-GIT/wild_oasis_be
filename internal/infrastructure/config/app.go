package config

import (
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/api/handlers"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/go-chi/chi/v5"
)

func Bootstrap() *chi.Mux {
	v := NewViper()
	l := NewLogger()
	db := NewDatabase(v, &l)

	// repository
	cabinRepository := repository.CabinRepository{}

	// services
	cloudinaryService := services.NewCloudinaryService(v)
	cabinService := services.NewCabinService(&cabinRepository, db, cloudinaryService)

	// handlers
	cabinHandler := handlers.NewCabinHandler(cabinService)
	cloudinaryHandler := handlers.NewCloudinaryHandler(cloudinaryService)

	router := handlers.NewRouter(cabinHandler, cloudinaryHandler)

	return router.GetRouter()
}
