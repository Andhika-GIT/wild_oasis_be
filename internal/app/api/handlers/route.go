package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	route        *chi.Mux
	CabinHandler *CabinHandler
}

func NewRouter(cabinHandler *CabinHandler) *Router {

	r := &Router{
		route:        chi.NewMux(),
		CabinHandler: cabinHandler,
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

		api.Get("/cabins", r.CabinHandler.FindAllCabins)
		api.Post("/cabins/seeds", r.CabinHandler.SeedsCabins)

	})

}

func (r *Router) GetRouter() *chi.Mux {
	return r.route
}
