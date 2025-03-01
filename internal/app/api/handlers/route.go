package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	route             *chi.Mux
	CabinHandler      *CabinHandler
	BookingHandler    *BookingHandler
	SettingHandler    *SettingHandler
	CloudinaryHandler *CloudinaryHandler
	AuthHandler       *AuthHandler
}

func NewRouter(cabinHandler *CabinHandler, bookingHandler *BookingHandler, settingHandler *SettingHandler, cloudinaryHandler *CloudinaryHandler, AuthHandler *AuthHandler) *Router {

	r := &Router{
		route:             chi.NewMux(),
		CabinHandler:      cabinHandler,
		BookingHandler:    bookingHandler,
		SettingHandler:    settingHandler,
		CloudinaryHandler: cloudinaryHandler,
		AuthHandler:       AuthHandler,
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
		api.Get("/cabins/{cabinId}", r.CabinHandler.FindCabinById)

		// bookings
		api.Get("/booking/booked-dates/cabin/{cabinId}", r.BookingHandler.GetBookedDatesByCabinId)

		// settings
		api.Get("/setting", r.SettingHandler.GetSetting)

		// auth
		api.Post("/auth/sign-in", r.AuthHandler.SignIn)
		api.Post("/auth/sign-up", r.AuthHandler.SignUp)

	})

}

func (r *Router) GetRouter() *chi.Mux {
	return r.route
}
