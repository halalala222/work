package usecase

import (
	"context"
	"go.uber.org/zap"
	"work/internal/domain"
	"work/internal/pkg/response"
)

type SalesUseCase struct {
	repo domain.SalesRepository
}

func NewSalesUseCase(repo domain.SalesRepository) domain.SalesUseCase {
	return &SalesUseCase{
		repo: repo,
	}
}

func (s *SalesUseCase) GetAllSales(ctx context.Context) ([]domain.Sales, error) {
	var (
		sales []domain.Sales
		err   error
	)
	if sales, err = s.repo.FetchAllSales(ctx); err != nil {
		zap.L().Error("fetch all sales error", zap.Error(err))
		return sales, response.ServerError
	}

	return sales, nil
}
