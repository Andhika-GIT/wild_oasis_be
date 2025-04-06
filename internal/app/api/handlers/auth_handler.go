package handlers

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/spf13/viper"
)

type AuthHandler struct {
	service *services.AuthService
	env     *viper.Viper
}

func NewAuthHandler(service *services.AuthService, viper *viper.Viper) *AuthHandler {
	return &AuthHandler{
		service: service,
		env:     viper,
	}
}

func (c *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.VerifyUser{}
	utils.ReadBodyRequest(r, bodyRequest)

	if bodyRequest.Email == "" || bodyRequest.Password == "" {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "All fields are required",
		})
		return
	}

	jwtToken, err := c.service.VerifyUser(r.Context(), *bodyRequest)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	utils.SetCookie(w, jwtToken, c.env.GetBool("IS_PRODUCTION"))

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully login",
		Data:    jwtToken,
	})
}

func (c *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.CreateUser{}
	utils.ReadBodyRequest(r, bodyRequest)

	if bodyRequest.Email == "" || bodyRequest.Password == "" {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "All fields are required",
		})
		return
	}

	isEmailExist := c.service.UserEmailExist(r.Context(), bodyRequest.Email)
	if isEmailExist {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Email already exists",
		})
		return
	}

	jwtToken, err := c.service.CreateNewUser(r.Context(), *bodyRequest)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	utils.SetCookie(w, jwtToken, c.env.GetBool("IS_PRODUCTION"))

	utils.SendResponse(w, http.StatusCreated, web.Response{
		Success: true,
		Code:    http.StatusCreated,
		Message: "Successfully created user",
		Data:    jwtToken,
	})
}
