package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SuccessWithData(ctx *fiber.Ctx, data any) error {
	return response(ctx, NewSuccessResult(data), http.StatusOK)
}

func SuccessWithoutData(ctx *fiber.Ctx) error {
	return response(ctx, SuccessWithoutDataResult, http.StatusOK)
}

func response(ctx *fiber.Ctx, data any, httpStatus int) error {
	return ctx.Status(httpStatus).JSON(data)
}
