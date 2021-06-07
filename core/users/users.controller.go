package users

import (
	"crud/core/users/dto"
	"crud/models"
	"crud/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	s *Service
}

func NewController(app fiber.Router) (c *Controller) {
	routes := app.Group("/users")
	c = &Controller{NewService()}
	routes.Get("/", c.getAll)
	routes.Post("/", c.create)
	routes.Get("/:id", c.getOne)
	routes.Patch("/:id", c.update)
	routes.Delete("/:id", c.delete)
	return

}
func (c *Controller) create(ctx *fiber.Ctx) error {
	user := new(dto.CreateUserDto)
	if err := ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "The body is not correct")
	}
	err := utils.ValidationStruct(user)
	if err != nil {
		return err
	}
	userCreated, err := c.s.create(user)
	if err != nil {
		return err
	}

	return ctx.JSON(models.ResponseSuccess{Data: userCreated, Message: "Creation successful "})
}
func (c *Controller) update(ctx *fiber.Ctx) error {

	userToUpdate := new(dto.UpdateUserDto)
	if err := ctx.BodyParser(userToUpdate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "The body is not correct")
	}
	id := ctx.Params("id")
	idInt, _ := strconv.ParseUint(id, 10, 64)
	user := c.s.update(userToUpdate, uint(idInt))
	return ctx.JSON(models.ResponseSuccess{Data: user, Message: "Updated successful"})

}
func (c *Controller) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.s.delete(id); err != nil {
		return err
	}
	return ctx.JSON(models.ResponseSuccess{Data: id, Message: "Deleted successful"})

}
func (c *Controller) getAll(ctx *fiber.Ctx) (err error) {
	users := c.s.getAll()
	err = ctx.JSON(models.ResponseSuccess{Data: users})
	return

}
func (c *Controller) getOne(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	idInt, _ := strconv.ParseUint(id, 10, 64)
	users := c.s.getOne(uint(idInt))
	err = ctx.JSON(models.ResponseSuccess{Data: users})
	return

}
