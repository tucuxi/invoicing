package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
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
		logger.New(),
		limiter.New(limiterConfig),
		idempotency.New(),
	)

	app.Post("/invoices", handlers.CreateInvoice())
	app.Post("/invoices/:id", handlers.UpdateInvoice())
	app.Get("/invoices/upcoming", handlers.RetrieveUpcomingInvoice()) // before GET "/invoices/:id"
	app.Get("/invoices/:id", handlers.RetrieveInvoice())
	app.Post("/invoices/:id/finalize", handlers.FinalizeInvoice())
	app.Delete("/invoices/:id", handlers.DeleteDraftInvoice())
	app.Post("/invoices/:id/mark_uncollectible", handlers.MarkInvoiceUncollectible())
	app.Post("/invoices/:id/pay", handlers.PayInvoice())
	app.Post("/invoices/:id/send", handlers.SendInvoice())
	app.Post("/invoices/:id/void", handlers.VoidInvoice())
	app.Post("/invoices/:id/lines/:line", handlers.UpdateLineItem())
	app.Get("/invoices/:id/lines", handlers.RetrieveLineItems())

	app.Listen(":3000")
}
