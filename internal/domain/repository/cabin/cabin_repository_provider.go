package repository

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"gorm.io/gorm"
)

type CabinRepositoryProvider struct{}

func NewCabinRepository() CabinRepository {
	return &CabinRepositoryProvider{}
}

func (r *CabinRepositoryProvider) Create(c context.Context, tx *gorm.DB, cabin entities.Cabin) error {
	err := tx.Create(cabin).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *CabinRepositoryProvider) FindAll(c context.Context, tx *gorm.DB, cabins []entities.Cabin) error {
	err := tx.Find(&cabins).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *CabinRepositoryProvider) FindById(c context.Context, tx *gorm.DB, cabinId int, cabin entities.Cabin) error {

	err := tx.Where("id = ?", cabinId).First(&cabin).Error

	if err != nil {
		return err
	}

	return nil

}

func (r *CabinRepositoryProvider) Delete(c context.Context, tx *gorm.DB, cabin entities.Cabin) error {
	err := tx.Delete(&cabin).Error

	if err != nil {
		return err
	}

	return nil
}
