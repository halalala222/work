package delivery

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"work/internal/domain"
	"work/internal/pkg/response"
)

type DrugDelivery struct {
	drugUseCase domain.DrugUseCase
}

func NewDrugDelivery(drugUseCase domain.DrugUseCase) *DrugDelivery {
	return &DrugDelivery{
		drugUseCase: drugUseCase,
	}

}

func (d *DrugDelivery) Create(c *fiber.Ctx) error {
	var drug *domain.Drug
	if err := c.BodyParser(&drug); err != nil {
		zap.L().Error("body parser error", zap.Error(err))
		return response.FormError
	}
	if err := d.drugUseCase.CreateOneDrug(c.Context(), drug); err != nil {
		return err
	}

	return response.SuccessWithoutData(c)
}

func (d *DrugDelivery) Update(c *fiber.Ctx) error {
	var drug *domain.Drug
	if err := c.BodyParser(&drug); err != nil {
		zap.L().Error("body parser error", zap.Error(err))
		return response.FormError
	}
	if err := d.drugUseCase.UpdateOneDrug(c.Context(), drug); err != nil {
		return err
	}

	return response.SuccessWithoutData(c)
}

func (d *DrugDelivery) GetAll(c *fiber.Ctx) error {
	var (
		drugs []domain.Drug
		err   error
	)
	if drugs, err = d.drugUseCase.GetAllDrugs(c.Context()); err != nil {
		return err
	}

	return response.SuccessWithData(c, drugs)
}

func (d *DrugDelivery) Delete(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		zap.L().Error("ParseInt error", zap.Error(err))
		return response.ServerError
	}

	if err = d.drugUseCase.DeleteOneDrug(c.Context(), id); err != nil {
		return err
	}

	return response.SuccessWithoutData(c)
}

func DrugRouter(router fiber.Router, handler *DrugDelivery) {
	drug := router.Group("/drug")
	drug.Get("", handler.GetAll)
	drug.Post("", handler.Create)
	drug.Put("", handler.Update)
	drug.Delete("/:id", handler.Delete)
}
