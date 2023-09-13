package repository

import (
	"context"
	"gorm.io/gorm"
	"work/internal/domain"
)

type DrugRepository struct {
	db *gorm.DB
}

func NewDrugRepo(db *gorm.DB) domain.DrugRepository {
	return &DrugRepository{
		db: db,
	}
}

func (d *DrugRepository) FetchAllDrug(ctx context.Context) ([]domain.Drug, error) {
	drugs := make([]domain.Drug, 0)
	err := d.db.WithContext(ctx).Model(&domain.Drug{}).Find(&drugs).Error
	return drugs, err
}

func (d *DrugRepository) CreateDrug(ctx context.Context, drug *domain.Drug) error {
	return d.db.WithContext(ctx).Model(&domain.Drug{}).Create(drug).Error
}

func (d *DrugRepository) UpdateDrug(ctx context.Context, drug *domain.Drug) error {
	return d.db.WithContext(ctx).Model(&domain.Drug{}).Where("id = ?", drug.Id).Updates(drug).Error
}

func (d *DrugRepository) DeleteDrug(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Delete(&domain.Drug{}, id).Error
}
