package handlers

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
)

type CabinController struct {
	service *services.CabinService
}

func NewCabinController(service *services.CabinService) *CabinController {
	return &CabinController{
		service: service,
	}
}

func (c *CabinController) FindAll(w http.ResponseWriter, r *http.Request) {

}
