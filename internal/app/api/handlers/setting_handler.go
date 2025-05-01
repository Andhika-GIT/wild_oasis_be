package handlers

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
)

type SettingHandler struct {
	service *services.SettingService
}

func NewSettingHandler(service *services.SettingService) *SettingHandler {
	return &SettingHandler{
		service: service,
	}
}

func (c *SettingHandler) GetSetting(w http.ResponseWriter, r *http.Request) {
	settingResponse, err := c.service.GetSetting(r.Context())

	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Sucessfully get setting",
		Data:    settingResponse,
	})
}
