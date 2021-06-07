package handlers

import (
	"crud/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	errorResponse := models.ResponseError{
		Code:       code,
		Method:     c.Method(),
		Path:       c.Path(),
		TimeStamps: time.Now(),
		Message:    err.Error(),
	}
	return c.Status(code).JSON(errorResponse)
}
