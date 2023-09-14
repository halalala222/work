package usecase

import (
	"context"
	"go.uber.org/zap"
	"work/internal/domain"
	"work/internal/pkg/response"
)

type DoctorUseCase struct {
	repo domain.DoctorRepository
}

func NewDoctorUseCase(repo domain.DoctorRepository) domain.DoctorUseCase {
	return &DoctorUseCase{
		repo: repo,
	}
}

func (d *DoctorUseCase) GetAllDoctors(ctx context.Context) ([]domain.Doctor, error) {
	var (
		doctors []domain.Doctor
		err     error
	)
	if doctors, err = d.repo.FetchAllDoctor(ctx); err != nil {
		zap.L().Error("fetch all doctor error", zap.Error(err))
		return doctors, response.ServerError
	}

	return doctors, nil
}
