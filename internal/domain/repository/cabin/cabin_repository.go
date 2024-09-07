package repository

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"
	"gorm.io/gorm"
)

type CabinRepository interface {
	Create(ctx context.Context, tx *gorm.DB, cabin entities.Cabin) error
	FindAll(ctx context.Context, tx *gorm.DB, cabins []entities.Cabin) error
	FindById(ctx context.Context, tx *gorm.DB, cabinId int, cabin entities.Cabin) error
	Delete(ctx context.Context, tx *gorm.DB, cabin entities.Cabin) error
}
