package domain

type Drug struct {
	Id          int64 `gorm:"id"`
	Name        string
	Measurement float64
	Type        string
}
