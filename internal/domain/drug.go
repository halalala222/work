package domain

import "context"

type Drug struct {
	Id          int64   `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Name        string  `json:"name"`
	Measurement float64 `json:"measurement"`
	Type        string  `json:"type"`
}

type DrugRepository interface {
	FetchAllDrug(ctx context.Context) ([]Drug, error)
	CreateDrug(ctx context.Context, drug *Drug) error
	UpdateDrug(ctx context.Context, drug *Drug) error
	DeleteDrug(ctx context.Context, id int64) error
}

type DrugUseCase interface {
	GetAllDrugs(ctx context.Context) ([]Drug, error)
	CreateOneDrug(ctx context.Context, drug *Drug) error
	UpdateOneDrug(ctx context.Context, drug *Drug) error
	DeleteOneDrug(ctx context.Context, id int64) error
}
