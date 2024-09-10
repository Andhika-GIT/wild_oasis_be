package handlers

import (
	"fmt"
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
)

type CabinHandler struct {
	service *services.CabinService
}

func NewCabinHandler(service *services.CabinService) *CabinHandler {
	return &CabinHandler{
		service: service,
	}
}

func (c *CabinHandler) SeedsCabins(w http.ResponseWriter, r *http.Request) {
	err := c.service.SeedCabins(r.Context())

	if err != nil {
		utils.SendResponse(w, web.Response{
			Code:    500,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, web.Response{
		Code:    200,
		Message: "Successfully seeds cabins",
	})
}

func (c *CabinHandler) FindAllCabins(w http.ResponseWriter, r *http.Request) {
	cabinResponse, err := c.service.FindAll(r.Context())

	if err != nil {

		utils.SendResponse(w, web.Response{
			Code:    500,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, web.Response{
		Code:    200,
		Message: "Successfully get all cabins",
		Data:    cabinResponse,
	})
}
