package utils

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ValidationStruct(item interface{}) error {
	err := validator.New().Struct(item)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("error in validation", err.Error())
			return &fiber.Error{
				Message: err.Error(),
				Code:    fiber.StatusBadRequest,
			}
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("The %v is %v but has been received %v", err.Field(), err.Tag(), err.Value())
			errors = append(errors, message)
			fmt.Println(message)
		}
		return &fiber.Error{
			Message: strings.Join(errors, ", "),
			Code:    fiber.StatusBadRequest,
		}
	}
	return nil
}
