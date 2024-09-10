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
	cabinService := services.NewCabinService(&cabinRepository, db)

	// handlers
	cabinHandler := handlers.NewCabinHandler(cabinService)

	router := handlers.NewRouter(cabinHandler)

	return router.GetRouter()
}
