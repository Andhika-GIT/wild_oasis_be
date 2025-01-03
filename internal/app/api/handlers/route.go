package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	route             *chi.Mux
	CabinHandler      *CabinHandler
	BookingHandler    *BookingHandler
	CloudinaryHandler *CloudinaryHandler
}

func NewRouter(cabinHandler *CabinHandler, bookingHandler *BookingHandler, cloudinaryHandler *CloudinaryHandler) *Router {

	r := &Router{
		route:             chi.NewMux(),
		CabinHandler:      cabinHandler,
		BookingHandler:    bookingHandler,
		CloudinaryHandler: cloudinaryHandler,
	}

	r.SetupRoute()

	return r
}

func (r *Router) SetupRoute() {

	// Group routes under /api prefix
	r.route.Route("/api", func(api chi.Router) {
		api.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from /api"))
		})

		// cloudinary
		api.Get("/cloudinary/image-asset", r.CloudinaryHandler.CheckImageAssets)
		api.Get("/cloudinary/get-URL", r.CloudinaryHandler.GetImagePublicUrl)

		// cabins
		api.Get("/cabins", r.CabinHandler.FindAllCabins)
		api.Post("/cabins/seeds", r.CabinHandler.SeedsCabins)
		api.Get("/cabins/{cabinId}", r.CabinHandler.FindCabinById)

		//booking
		api.Get("/booking/find-booked-dates-by-cabin/{cabinId}", r.BookingHandler.GetBookedDatesByCabinId)

	})

}

func (r *Router) GetRouter() *chi.Mux {
	return r.route
}
