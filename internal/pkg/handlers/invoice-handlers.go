package handlers

import (
	"log/slog"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tucuxi/invoicing/internal/pkg/repository"
	"github.com/tucuxi/invoicing/pkg/invoice"
	"github.com/valyala/fasthttp"
)

func CreateInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		i := new(invoice.Invoice)

		if err := c.BodyParser(i); err != nil {
			slog.Error("CreateInvoice", "parse body error", err)
			c.SendStatus(fasthttp.StatusInternalServerError)
			return err
		}
		if slices.Index(invoice.InvoiceTypes, i.Type) == -1 {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		if i.Recipient == "" {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		i.ID = invoice.NewInvoiceID()
		i.Status = invoice.Draft
		i.DraftedAt = time.Now().Unix()
		if err := r.CreateInvoice(i); err != nil {
			c.SendStatus(fasthttp.StatusInternalServerError)
			return err
		}
		return c.JSON(i)
	}
}

func UpdateInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("update invoice " + id)
	}
}

func RetrieveUpcomingInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := c.FormValue("customer")
		return c.SendString("retrieve upcoming invoice for customer " + customer)
	}
}

func RetrieveInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := r.FindInvoice(id)
		if err == repository.ErrorInvoiceNotFound {
			c.SendStatus(fasthttp.StatusNotFound)
			return err
		}
		if err != nil {
			c.SendStatus(fasthttp.StatusInternalServerError)
			return err
		}
		return c.JSON(i)
	}
}

func UpdateLineItem(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		line := c.Params("line")
		return c.SendString("update line " + line + " of invoice " + id)
	}
}

func RetrieveLineItems(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("retrieve line items for invoice " + id)
	}
}

func DeleteDraftInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("delete draft invoice " + id)
	}
}

func FinalizeInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("finalize invoice " + id)
	}
}

func MarkInvoiceUncollectible(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("mark invoice " + id + " uncollectible")
	}
}

func PayInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("pay invoice " + id)
	}
}

func SendInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("send invoice " + id)
	}
}

func VoidInvoice(r *repository.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("void invoice " + id)
	}
}
