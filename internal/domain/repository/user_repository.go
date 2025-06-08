package repository

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) Create(c context.Context, tx *gorm.DB, user *entities.User) error {
	err := tx.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateNationality(c context.Context, tx *gorm.DB, user *entities.User, updatedData *web.UpdateUserNationality) error {
	user.NationalID = updatedData.NationalID
	user.Nationality = updatedData.Nationality
	user.CountryFlag = updatedData.CountryFlag

	err := tx.Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindById(c context.Context, tx *gorm.DB, userId int, user *entities.User) error {
	err := tx.Where("id = ?", userId).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(c context.Context, tx *gorm.DB, userEmail string, user *entities.User) error {
	err := tx.Where("email = ?", userEmail).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByPassword(c context.Context, tx *gorm.DB, userPassword string, user *entities.User) error {
	err := tx.Where("password = ?", userPassword).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}
