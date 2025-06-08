package handlers

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/services"
	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	utils "github.com/Andhika-GIT/wild_oasis_be/pkg/web"
	"github.com/go-chi/jwtauth/v5"
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

	utils.SetCookie(w, "access_token", jwtToken, c.env.GetBool("IS_PRODUCTION"))

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully login",
		Data:    nil,
	})
}

func (c *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.CreateUser{}
	utils.ReadBodyRequest(r, bodyRequest)

	if bodyRequest.Email == "" || bodyRequest.Password == "" || bodyRequest.Fullname == "" {
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

	utils.SetCookie(w, "access_token", jwtToken, c.env.GetBool("IS_PRODUCTION"))

	utils.SendResponse(w, http.StatusCreated, web.Response{
		Success: true,
		Code:    http.StatusCreated,
		Message: "Successfully created user",
		Data:    nil,
	})
}

func (c *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	_, _, err := jwtauth.FromContext(r.Context())

	if err != nil {
		utils.SendResponse(w, http.StatusUnauthorized, web.Response{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized - No token",
		})
		return
	}

	utils.ClearCookie(w, "access_token", c.env.GetBool("IS_PRODUCTION"))

	utils.SendResponse(w, http.StatusCreated, web.Response{
		Success: true,
		Code:    http.StatusCreated,
		Message: "Successfully sign out",
		Data:    nil,
	})

}

func (c *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.GetUserIDFromToken(r)

	if err != nil {
		utils.SendResponse(w, http.StatusUnauthorized, web.Response{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	userData, err := c.service.FindCurrentUser(r.Context(), userID)

	if err != nil {
		utils.SendResponse(w, http.StatusNotFound, web.Response{
			Success: false,
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	utils.SendResponse(w, http.StatusCreated, web.Response{
		Success: true,
		Code:    http.StatusCreated,
		Message: "Successfully get user",
		Data:    userData,
	})
}

func (c *AuthHandler) UpdateCurrentUserNationality(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &web.UpdateUserNationality{}
	utils.ReadBodyRequest(r, bodyRequest)

	if bodyRequest.NationalID == "" || bodyRequest.Nationality == "" || bodyRequest.CountryFlag == "" {
		utils.SendResponse(w, http.StatusBadRequest, web.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "All fields are required",
		})
		return
	}

	userID, err := utils.GetUserIDFromToken(r)

	if err != nil {
		utils.SendResponse(w, http.StatusUnauthorized, web.Response{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	err = c.service.UpdateUserNationality(r.Context(), userID, *bodyRequest)

	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, web.Response{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	utils.SendResponse(w, http.StatusOK, web.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "Successfully update user",
	})
}
