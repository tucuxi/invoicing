package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tucuxi/invoices/internal/pkg/invoice"
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

	app.Post("/invoices", invoice.CreateInvoice())
	app.Post("/invoices/:id", invoice.UpdateInvoice())
	app.Get("/invoices/upcoming", invoice.RetrieveUpcomingInvoice()) // before GET "/invoices/:id"
	app.Get("/invoices/:id", invoice.RetrieveInvoice())
	app.Post("/invoices/:id/finalize", invoice.FinalizeInvoice())
	app.Delete("/invoices/:id", invoice.DeleteDraftInvoice())
	app.Post("/invoices/:id/mark_uncollectible", invoice.MarkInvoiceUncollectible())
	app.Post("/invoices/:id/pay", invoice.PayInvoice())
	app.Post("/invoices/:id/send", invoice.SendInvoice())
	app.Post("/invoices/:id/void", invoice.VoidInvoice())
	app.Post("/invoices/:id/lines/:line", invoice.UpdateLineItem())
	app.Get("/invoices/:id/lines", invoice.RetrieveLineItems())

	app.Listen(":3000")
}
