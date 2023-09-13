package usecase

import (
	"context"
	"go.uber.org/zap"
	"work/internal/domain"
	"work/internal/pkg/response"
)

type DrugUseCase struct {
	repo domain.DrugRepository
}

func NewDrugUseCase(repo domain.DrugRepository) domain.DrugUseCase {
	return &DrugUseCase{
		repo: repo,
	}
}

func (d *DrugUseCase) GetAllDrugs(ctx context.Context) ([]domain.Drug, error) {
	var (
		drugs []domain.Drug
		err   error
	)

	if drugs, err = d.repo.FetchAllDrug(ctx); err != nil {
		zap.L().Error("fetch all drug error", zap.Error(err))
		return drugs, response.ServerError
	}

	return drugs, nil
}

func (d *DrugUseCase) CreateOneDrug(ctx context.Context, drug *domain.Drug) error {
	if err := d.repo.CreateDrug(ctx, drug); err != nil {
		zap.L().Error("create drug error", zap.Error(err))
		return response.ServerError
	}

	return nil
}

func (d *DrugUseCase) UpdateOneDrug(ctx context.Context, drug *domain.Drug) error {
	if err := d.repo.UpdateDrug(ctx, drug); err != nil {
		zap.L().Error("update drug error", zap.Error(err))
		return response.ServerError
	}

	return nil
}

func (d *DrugUseCase) DeleteOneDrug(ctx context.Context, id int64) error {
	if err := d.repo.DeleteDrug(ctx, id); err != nil {
		zap.L().Error("delete drug error", zap.Error(err))
		return response.ServerError
	}

	return nil
}
