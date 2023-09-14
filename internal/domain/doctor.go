package domain

import "context"

type Doctor struct {
	Id         int64  `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Position   string `json:"position"`
}

type DoctorRepository interface {
	FetchAllDoctor(ctx context.Context) ([]Doctor, error)
}

type DoctorUseCase interface {
	GetAllDoctors(ctx context.Context) ([]Doctor, error)
}
