package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tucuxi/invoices/internal/handlers"
)

func main() {
	app := fiber.New()

	limiterConfig := limiter.Config{
		Max:        5,
		Expiration: 1 * time.Second,
	}

	app.Use(
		limiter.New(limiterConfig),
		logger.New(),
	)

	app.Post("/invoices", handlers.CreateInvoice())
	app.Post("/invoices/:invoice", handlers.UpdateInvoice())
	app.Post("/invoices/:invoice/lines/:line", handlers.UpdateLineItem())
	app.Get("/invoices/:invoice", handlers.RetrieveInvoice())

	app.Listen(":3000")
}
