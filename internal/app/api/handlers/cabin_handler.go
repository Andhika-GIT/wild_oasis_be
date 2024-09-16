package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/go-chi/chi/v5"
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
		utils.SendResponse(w, 500, web.Response{
			Code:    500,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, 200, web.Response{
		Code:    200,
		Message: "Successfully seeds cabins",
	})
}

func (c *CabinHandler) FindAllCabins(w http.ResponseWriter, r *http.Request) {
	cabinResponse, err := c.service.FindAll(r.Context())

	if err != nil {

		utils.SendResponse(w, 500, web.ErrorResponse{
			Code:    500,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, 200, web.Response{
		Code:    200,
		Message: "Successfully get all cabins",
		Data:    cabinResponse,
	})
}

func (c *CabinHandler) FindCabinById(w http.ResponseWriter, r *http.Request) {
	cabinId := chi.URLParam(r, "cabinId")

	if cabinId == "" {
		utils.SendResponse(w, 400, web.ErrorResponse{
			Code:    400,
			Message: "Cabin ID is required",
		})

		return
	}

	id, err := strconv.Atoi(cabinId)

	if err != nil {
		utils.SendResponse(w, 500, web.ErrorResponse{
			Code:    500,
			Message: "Something went wrong",
		})

		return
	}

	cabinResponse, err := c.service.FindById(r.Context(), id)

	if err != nil {
		utils.SendResponse(w, 400, web.ErrorResponse{
			Code:    404,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, 200, web.Response{
		Code:    200,
		Message: "Sucessfully get cabin",
		Data:    cabinResponse,
	})
}
