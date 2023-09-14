package repository

import (
	"context"
	"gorm.io/gorm"
	"work/internal/domain"
)

type DoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) domain.DoctorRepository {
	return &DoctorRepository{
		db: db,
	}
}

func (d *DoctorRepository) FetchAllDoctor(ctx context.Context) ([]domain.Doctor, error) {
	var doctors []domain.Doctor
	err := d.db.WithContext(ctx).Model(&domain.Doctor{}).Find(&doctors).Error
	return doctors, err
}
