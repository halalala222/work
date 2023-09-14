package domain

import "context"

type Sales struct {
	Id            int64   `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Name          string  `json:"name"`
	PurchasePrice float64 `json:"purchasePrice"`
	SellingPrice  float64 `json:"sellingPrice"`
}

type SalesRepository interface {
	FetchAllSales(ctx context.Context) ([]Sales, error)
}

type SalesUseCase interface {
	GetAllSales(ctx context.Context) ([]Sales, error)
}
