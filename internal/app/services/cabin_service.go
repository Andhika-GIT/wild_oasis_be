package services

import (
	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/repository"
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
