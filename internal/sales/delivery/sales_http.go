package delivery

import (
	"github.com/gofiber/fiber/v2"
	"work/internal/domain"
	"work/internal/pkg/response"
)

type SalesDelivery struct {
	salesUseCase domain.SalesUseCase
}

func NewSalesDelivery(salesUseCase domain.SalesUseCase) *SalesDelivery {
	return &SalesDelivery{
		salesUseCase: salesUseCase,
	}
}

func (s *SalesDelivery) Get(c *fiber.Ctx) error {
	var (
		sales []domain.Sales
		err   error
	)
	if sales, err = s.salesUseCase.GetAllSales(c.Context()); err != nil {
		return err
	}

	return response.SuccessWithData(c, sales)
}

func SalesRouter(router fiber.Router, handler *SalesDelivery) {
	sales := router.Group("/sales")
	{
		sales.Get("", handler.Get)
	}
}
