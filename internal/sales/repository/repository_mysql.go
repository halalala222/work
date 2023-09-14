package repository

import (
	"context"
	"gorm.io/gorm"
	"work/internal/domain"
)

type SalesRepository struct {
	db *gorm.DB
}

func NewSalesRepo(db *gorm.DB) domain.SalesRepository {
	return &SalesRepository{
		db: db,
	}
}

func (s *SalesRepository) FetchAllSales(ctx context.Context) ([]domain.Sales, error) {
	var sales []domain.Sales
	err := s.db.WithContext(ctx).Model(&domain.Sales{}).Find(&sales).Error
	return sales, err
}
