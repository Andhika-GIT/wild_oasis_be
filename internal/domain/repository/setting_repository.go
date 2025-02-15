package repository

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"gorm.io/gorm"
)

type SettingRepository struct{}

func (r *SettingRepository) FindSetting(c context.Context, tx *gorm.DB, setting *entities.Settings) error {
	err := tx.Find(&setting).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *SettingRepository) Create(c context.Context, tx *gorm.DB, setting *entities.Settings) error {
	err := tx.Create(&setting).Error

	if err != nil {
		return err
	}

	return nil
}
