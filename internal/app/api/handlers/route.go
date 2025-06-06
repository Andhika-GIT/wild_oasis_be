package handlers

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/api/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

type Router struct {
	route             *chi.Mux
	CabinHandler      *CabinHandler
	BookingHandler    *BookingHandler
	SettingHandler    *SettingHandler
	CloudinaryHandler *CloudinaryHandler
	AuthHandler       *AuthHandler
	env               *viper.Viper
}

func NewRouter(cabinHandler *CabinHandler, bookingHandler *BookingHandler, settingHandler *SettingHandler, cloudinaryHandler *CloudinaryHandler, AuthHandler *AuthHandler, env *viper.Viper) *Router {

	r := &Router{
		route:             chi.NewMux(),
		CabinHandler:      cabinHandler,
		BookingHandler:    bookingHandler,
		SettingHandler:    settingHandler,
		CloudinaryHandler: cloudinaryHandler,
		AuthHandler:       AuthHandler,
		env:               env,
	}

	r.SetupRoute()

	return r
}

func (r *Router) SetupRoute() {

	jwt_secret := r.env.GetString("JWT_SECRET")

	r.route.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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

		api.Group(func(protected chi.Router) {

			protected.Use(middleware.AuthMiddleware(jwt_secret))
			protected.Get("/auth/sign-out", r.AuthHandler.SignOut)
			protected.Get("/auth/me", r.AuthHandler.GetCurrentUser)
			protected.Put("/auth/update-nationality", r.AuthHandler.UpdateCurrentUserNationality)
			protected.Get("/booking/me", r.BookingHandler.GetAllUserBookings)
			protected.Post("/booking/create", r.BookingHandler.CreateUserBooking)

			protected.Group(func(protected_booking chi.Router) {
				protected_booking.Use(middleware.UserBookingMiddleware(r.BookingHandler.BookingService))
				protected_booking.Get("/booking/me/{bookingId}", r.BookingHandler.GetSpesificUserBooking)
				protected_booking.Delete("/booking/me/{bookingId}", r.BookingHandler.DeleteCurrentUserBooking)
				protected_booking.Put("/booking/me/{bookingId}", r.BookingHandler.UpdateUserBooking)
			})

		})

	})

}

func (r *Router) GetRouter() *chi.Mux {
	return r.route
}
