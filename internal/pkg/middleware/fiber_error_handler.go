package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"work/internal/consts"
	"work/internal/pkg/response"
)

var CustomErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	httpStatus := fiber.StatusInternalServerError
	result := response.UnknownErrorResult

	var e *response.ErrorResult
	if errors.As(err, &e) {
		httpStatus = response.GetErrorHTTPStatus(e)
		result = response.GetErrorResult(e)
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		httpStatus = fiberError.Code
		result = response.NewErrorResult(consts.UnknownError, fiberError.Message)
	}

	// Return status code with error message
	return c.Status(httpStatus).JSON(result)
}
