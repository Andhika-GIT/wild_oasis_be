package handlers

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (c *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.VerifyUser{}

	utils.ReadBodyRequest(r, bodyRequest)

	if bodyRequest.Email == "" || bodyRequest.Password == "" {
		utils.SendResponse(w, http.StatusBadRequest, web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "All fields are required",
		})
		return
	}

	err := c.service.VerifyUser(r.Context(), *bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Successfully login",
		Data:    "",
	})

}

func (c *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.CreateUser{}

	utils.ReadBodyRequest(r, bodyRequest)

	if bodyRequest.Email == "" || bodyRequest.Password == "" {
		utils.SendResponse(w, http.StatusBadRequest, web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "All fields are required",
		})
		return
	}

	isEmailExist := c.service.UserEmailExist(r.Context(), bodyRequest.Email)

	if isEmailExist {
		utils.SendResponse(w, http.StatusBadRequest, web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already exist",
		})

		return
	}

	err := c.service.CreateNewUser(r.Context(), *bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

		return
	}

	utils.SendResponse(w, http.StatusCreated, web.Response{
		Code:    http.StatusOK,
		Message: "Successfully created user",
		Data:    "",
	})
}

// func (c *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
// 	bodyRequest := &web.CreateUser{}

// 	utils.ReadBodyRequest(r, bodyRequest)
// }
