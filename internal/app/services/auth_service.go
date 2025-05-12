package services

import (
	"context"
	"fmt"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository *repository.UserRepository
	DB         *gorm.DB
	env        *viper.Viper
}

func NewAuthService(repository *repository.UserRepository, DB *gorm.DB, viper *viper.Viper) *AuthService {
	return &AuthService{
		repository: repository,
		DB:         DB,
		env:        viper,
	}
}

func HashPassword(stringPassword string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(stringPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func checkPassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))

	return err == nil
}

func (s *AuthService) GenerateJwtToken(userId int) (string, error) {
	var tokenAuth = jwtauth.New("HS256", []byte(s.env.GetString("JWT_SECRET")), nil)

	_, tokenJwtString, err := tokenAuth.Encode(map[string]interface{}{"user_id": userId})

	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return tokenJwtString, nil
}

func (s *AuthService) UserEmailExist(c context.Context, userEmail string) bool {
	var user entities.User
	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	err := s.repository.FindByEmail(c, tx, userEmail, &user)

	if err != nil {
		return false
	}

	return true
}

func (s *AuthService) VerifyUser(c context.Context, inputUser web.VerifyUser) (string, error) {

	isEmailExist := s.UserEmailExist(c, inputUser.Email)

	if !isEmailExist {
		return "", fmt.Errorf("wrong credentials")
	}

	var user entities.User

	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	err := s.repository.FindByEmail(c, tx, inputUser.Email, &user)

	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}

	isPasswordCorrent := checkPassword(user.Password, inputUser.Password)

	if !isPasswordCorrent {
		return "", fmt.Errorf("wrong credentials")
	}

	jwtToken, err := s.GenerateJwtToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return jwtToken, tx.Commit().Error
}

func (s *AuthService) CreateNewUser(c context.Context, userData web.CreateUser) (string, error) {

	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	hashedPassword, err := HashPassword(userData.Password)

	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}

	user := entities.User{
		Email:    userData.Email,
		FullName: userData.Fullname,
		Password: hashedPassword,
	}

	err = s.repository.Create(c, tx, &user)

	if err != nil {
		return "", fmt.Errorf("error creating user : %v", err)
	}

	jwtToken, err := s.GenerateJwtToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return jwtToken, tx.Commit().Error
}

func (s *AuthService) FindCurrentUser(c context.Context, userID int) (web.UserResponse, error) {
	var user entities.User
	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	err := s.repository.FindById(c, tx, userID, &user)

	if err != nil {
		return web.UserResponse{}, fmt.Errorf("user not found")
	}

	return web.UserResponse{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		NationalID:  user.NationalID,
		Nationality: user.Nationality,
	}, tx.Commit().Error

}

func (s *AuthService) UpdateUserNationality(c context.Context, userID int, userData web.UpdateUserNationality) error {
	var user entities.User
	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	err := s.repository.FindById(c, tx, userID, &user)

	if err != nil {
		return fmt.Errorf("user not found")
	}

	err = s.repository.UpdateNationality(c, tx, &user, &userData)

	if err != nil {
		return fmt.Errorf("error updating user: %v", err.Error())
	}

	return tx.Commit().Error
}
