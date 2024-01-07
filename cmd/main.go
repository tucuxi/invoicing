package main

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/tucuxi/invoicing/internal/pkg/handlers"
	"github.com/tucuxi/invoicing/internal/pkg/persistence"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	limiterConfig := limiter.Config{
		Max:        5,
		Expiration: 1 * time.Second,
	}

	app.Use(
		logger.New(),
		limiter.New(limiterConfig),
		idempotency.New(),
	)

	r := persistence.NewInvoiceRepository()

	app.Post("/invoices", handlers.CreateInvoice(r))
	app.Post("/invoices/:id", handlers.UpdateInvoice(r))
	app.Get("/invoices/upcoming", handlers.RetrieveUpcomingInvoice(r)) // before GET "/invoices/:id"
	app.Get("/invoices/:id", handlers.RetrieveInvoice(r))
	app.Post("/invoices/:id/finalize", handlers.FinalizeInvoice(r))
	app.Delete("/invoices/:id", handlers.DeleteDraftInvoice(r))
	app.Post("/invoices/:id/mark_uncollectible", handlers.MarkInvoiceUncollectible(r))
	app.Post("/invoices/:id/pay", handlers.PayInvoice(r))
	app.Post("/invoices/:id/send", handlers.SendInvoice(r))
	app.Post("/invoices/:id/void", handlers.VoidInvoice(r))
	app.Post("/invoices/:id/lines/:line", handlers.UpdateLineItem(r))
	app.Get("/invoices/:id/lines", handlers.RetrieveLineItems(r))

	if err := app.Listen(":3000"); err != nil {
		slog.Error("Listen", err)
	}
}
