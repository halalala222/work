package delivery

import (
	"github.com/gofiber/fiber/v2"
	"work/internal/domain"
	"work/internal/pkg/response"
)

type DoctorDelivery struct {
	doctorUseCase domain.DoctorUseCase
}

func NewDoctorDelivery(doctorUseCase domain.DoctorUseCase) *DoctorDelivery {
	return &DoctorDelivery{
		doctorUseCase: doctorUseCase,
	}
}

func (d *DoctorDelivery) Get(c *fiber.Ctx) error {
	var (
		doctors []domain.Doctor
		err     error
	)
	if doctors, err = d.doctorUseCase.GetAllDoctors(c.Context()); err != nil {
		return err
	}

	return response.SuccessWithData(c, doctors)
}

func DoctorRouter(router fiber.Router, handler *DoctorDelivery) {
	doctor := router.Group("/doctor")
	{
		doctor.Get("", handler.Get)
	}
}
