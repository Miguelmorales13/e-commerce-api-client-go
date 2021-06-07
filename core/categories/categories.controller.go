package categories

import (
	"crud/models"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	s *Service
}

func NewController(app fiber.Router) (c *Controller) {
	routes := app.Group("/categories")
	c = &Controller{NewService()}
	routes.Get("/", c.getAll)
	routes.Post("/", c.create)
	routes.Post("/", c.update)
	return

}
func (c *Controller) create(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
func (c *Controller) update(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)

}
func (c *Controller) delete(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)

}
func (c *Controller) getAll(ctx *fiber.Ctx) (err error) {
	categories := c.s.getAll()
	err = ctx.JSON(models.ResponseSuccess{Data: categories})
	return

}
