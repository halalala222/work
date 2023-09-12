package domain

type Sales struct {
	Id            int64 `gorm:"id"`
	Name          string
	PurchasePrice float64
	SellingPrice  float64
}
