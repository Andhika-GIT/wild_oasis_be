package handlers

import (
	"fmt"
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
)

type CloudinaryHandler struct {
	service *services.CloudinaryService
}

func NewCloudinaryHandler(s *services.CloudinaryService) *CloudinaryHandler {
	return &CloudinaryHandler{
		service: s,
	}
}

func (c *CloudinaryHandler) CheckImageAssets(w http.ResponseWriter, r *http.Request) {
	publicID := r.URL.Query().Get("publicID")

	if publicID == "" {
		utils.SendResponse(w, 400, web.ErrorResponse{
			Code:    400,
			Message: "publicID is required",
		})
		return
	}

	res, err := c.service.CheckImageAssets(r.Context(), publicID)

	if err != nil {
		utils.SendResponse(w, 400, web.ErrorResponse{
			Code:    400,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, 200, web.Response{
		Code:    200,
		Message: "Sucessfully get image asset",
		Data:    res,
	})
}

func (c *CloudinaryHandler) GetImagePublicUrl(w http.ResponseWriter, r *http.Request) {
	publicID := r.URL.Query().Get("publicID")

	if publicID == "" {
		utils.SendResponse(w, 400, web.ErrorResponse{
			Code:    400,
			Message: "publicID is required",
		})
		return
	}

	imageURL, err := c.service.GetImagePublicUrl(r.Context(), publicID)

	if err != nil {
		utils.SendResponse(w, 400, web.ErrorResponse{
			Code:    400,
			Message: fmt.Sprintf("error, %s", err),
		})

		return
	}

	utils.SendResponse(w, 200, web.Response{
		Code:    200,
		Message: "Sucessfully get image URL",
		Data:    imageURL,
	})
}
