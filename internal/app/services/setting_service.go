package services

import (
	"context"
	"fmt"

	"github.com/Andhika-GIT/wild_oasis_be/internal/app/web"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
	"github.com/Andhika-GIT/wild_oasis_be/pkg/file"
	"gorm.io/gorm"
)

type SettingService struct {
	repository *repository.SettingRepository
	DB         *gorm.DB
}

func NewSettingService(repository *repository.SettingRepository, DB *gorm.DB) *SettingService {
	return &SettingService{
		repository: repository,
		DB:         DB,
	}
}

func (s *SettingService) GetSetting(c context.Context) (web.SettingResponse, error) {
	var setting entities.Settings

	tx := s.DB.WithContext(c).Begin()

	defer tx.Rollback()

	err := s.repository.FindSetting(c, tx, &setting)

	if err != nil {
		return web.SettingResponse{}, fmt.Errorf("error when finding settings : %v", err)
	}

	settingResponse := web.ToSettingResponse(setting)

	return settingResponse, tx.Commit().Error
}

func (s *SettingService) SeedSetting(c context.Context) error {

	tx := s.DB.WithContext(c).Begin()

	// rollback after all function done
	defer tx.Rollback()

	// read file from json
	settings, err := file.LoadFromJsonFile[[]entities.Settings]("./data/settings.json")
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	//  reset all data first first
	err = tx.Exec("DELETE from settings").Error
	if err != nil {
		return fmt.Errorf("error when deleting all settings : %v", err)
	}

	err = tx.Exec("TRUNCATE TABLE settings RESTART IDENTITY CASCADE").Error
	if err != nil {
		return fmt.Errorf("error when truncating table: %v", err)
	}

	for _, setting := range settings {
		err = s.repository.Create(c, tx, &setting)

		if err != nil {
			return fmt.Errorf("error create setting : %v", err)
		}
	}

	return tx.Commit().Error

}
