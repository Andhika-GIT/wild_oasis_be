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

type CabinService struct {
	repository *repository.CabinRepository
	DB         *gorm.DB
}

func NewCabinService(repository *repository.CabinRepository, DB *gorm.DB) *CabinService {
	return &CabinService{
		repository: repository,
		DB:         DB,
	}
}

func (s *CabinService) FindAll(c context.Context) ([]web.CabinResponse, error) {
	var cabins []entities.Cabin

	tx := s.DB.WithContext(c).Begin()

	// rollback after all function done
	defer tx.Rollback()

	err := s.repository.FindAll(c, tx, &cabins)

	if err != nil {
		return []web.CabinResponse{}, fmt.Errorf("error find all cabins: %v", err)
	}

	// format cabins for response
	cabinResponse := web.ToCabinResponses(cabins)

	return cabinResponse, tx.Commit().Error

}

func (s *CabinService) SeedCabins(c context.Context) error {

	tx := s.DB.WithContext(c).Begin()

	// rollback after all function done
	defer tx.Rollback()

	// read file from json
	cabins, err := file.LoadFromJsonFile[[]entities.Cabin]("./data/cabins.json")
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	//  reset all data first first
	err = tx.Exec("DELETE from cabins").Error
	if err != nil {
		return fmt.Errorf("error when deleting all cabins : %v", err)
	}

	for _, cabin := range cabins {
		err = s.repository.Create(c, tx, &cabin)

		if err != nil {
			return fmt.Errorf("error create cabin : %v", err)
		}
	}

	return tx.Commit().Error

}
