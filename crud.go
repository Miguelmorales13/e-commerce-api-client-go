package main

import (
	"crud/configurations"
	"crud/core/categories"
	"crud/core/products"
	"crud/core/users"
	"crud/handlers"
	"crud/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"time"
)

func main() {
	configurations.Config()
	providers.InitDbPostgres()

	app := fiber.New(
		fiber.Config{
			ErrorHandler: handlers.ErrorHandler,
		},
	)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(recover2.New())
	app.Use(requestid.New())
	app.Use(logger.New(
		logger.Config{
			Next:         nil,
			Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
			TimeFormat:   "15:04:05",
			TimeZone:     "Local",
			TimeInterval: 500 * time.Millisecond,
			Output:       os.Stderr},
	))
	api := app.Group("/api")
	users.NewController(api)
	products.NewController(api)
	categories.NewController(api)
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
