package repository

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"gorm.io/gorm"
)

type CabinRepository struct{}

func (r *CabinRepository) Create(c context.Context, tx *gorm.DB, cabin *entities.Cabin) error {
	err := tx.Create(&cabin).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CabinRepository) FindAll(c context.Context, tx *gorm.DB, cabins *[]entities.Cabin) error {
	err := tx.Find(&cabins).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CabinRepository) FindById(c context.Context, tx *gorm.DB, cabinId int, cabin *entities.Cabin) error {
	err := tx.Where("id = ?", cabinId).First(&cabin).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CabinRepository) Delete(c context.Context, tx *gorm.DB, cabin *entities.Cabin) error {
	err := tx.Delete(&cabin).Error
	if err != nil {
		return err
	}
	return nil
}
