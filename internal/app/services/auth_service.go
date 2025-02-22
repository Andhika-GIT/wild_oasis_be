package services

import (
	"context"
	"fmt"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository *repository.UserRepository
	DB         *gorm.DB
}

func NewAuthService(repository *repository.UserRepository, DB *gorm.DB) *AuthService {
	return &AuthService{
		repository: repository,
		DB:         DB,
	}
}

func HashPassword(stringPassword string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(stringPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
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

func (s *AuthService) CreateNewUser(c context.Context, userData web.CreateUser) error {

	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	hashedPassword, err := HashPassword(userData.Password)

	if err != nil {
		return fmt.Errorf("something went wrong")
	}

	user := entities.User{
		Email:    userData.Email,
		Password: hashedPassword,
	}

	err = s.repository.Create(c, tx, &user)

	if err != nil {
		return fmt.Errorf("error creating user : %v", err)
	}

	return tx.Commit().Error
}
