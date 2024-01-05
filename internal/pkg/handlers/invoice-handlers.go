package handlers

import (
	"errors"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tucuxi/invoicing/internal/pkg/persistence"
	"github.com/tucuxi/invoicing/pkg/invoice"
	"github.com/valyala/fasthttp"
)

type parameters struct {
	PaidOutOfBand bool  `json:"paid_out_of_band" form:"paid_out_of_band"`
	AmountPaid    int64 `json:"amount_paid" form:"amount_paid"`
}

func CreateInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		i := new(invoice.Invoice)

		if c.BodyParser(i) != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		if slices.Index(invoice.InvoiceTypes, i.Type) == -1 {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		if i.Recipient == "" {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		i.ID = invoice.NewInvoiceID()
		i.Status = invoice.StatusDraft
		i.DraftedAt = time.Now().Unix()
		if err := r.CreateInvoice(i); err != nil {
			return err
		}
		return c.JSON(i)
	}
}

func UpdateInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("update invoice " + id)
	}
}

func RetrieveUpcomingInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := c.FormValue("customer")
		return c.SendString("retrieve upcoming invoice for customer " + customer)
	}
}

func RetrieveInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := r.FindInvoice(id)
		switch {
		case err == nil:
			return c.JSON(i)
		case errors.Is(err, persistence.ErrorInvoiceNotFound):
			return c.SendStatus(fasthttp.StatusNotFound)
		default:
			return err
		}
	}
}

func UpdateLineItem(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		line := c.Params("line")
		return c.SendString("update line " + line + " of invoice " + id)
	}
}

func RetrieveLineItems(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("retrieve line items for invoice " + id)
	}
}

func DeleteDraftInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := r.DeleteDraftInvoice(id)
		switch {
		case errors.Is(err, persistence.ErrorInvoiceNotFound):
			return c.SendStatus(fasthttp.StatusNotFound)
		case errors.Is(err, persistence.ErrorDeletionNotAllowed):
			return c.SendStatus(fasthttp.StatusBadRequest)
		default:
			return err
		}
	}
}

func FinalizeInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("finalize invoice " + id)
	}
}

func MarkInvoiceUncollectible(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		switch i, err := r.FindInvoice(id); err {
		case nil:
			i.Status = invoice.StatusUncollectible
			i.MarkedUncollectibleAt = time.Now().Unix()
			return c.JSON(i)
		case persistence.ErrorInvoiceNotFound:
			return c.SendStatus(fasthttp.StatusNotFound)
		default:
			return err
		}
	}
}

func PayInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := new(parameters)

		if c.BodyParser(p) != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		if !p.PaidOutOfBand {
			return c.SendStatus(fasthttp.StatusNotImplemented)
		}

		id := c.Params("id")

		i, err := r.FindInvoice(id)
		if errors.Is(err, persistence.ErrorInvoiceNotFound) {
			return c.SendStatus(fasthttp.StatusNotFound)
		}
		if err != nil {
			return err
		}

		i.Status = invoice.StatusPaid
		i.PaidAt = time.Now().Unix()
		i.PaidOutOfBand = p.PaidOutOfBand
		i.AmountPaid = p.AmountPaid

		if err = r.UpdateInvoice(i); err != nil {
			return err
		}
		return c.JSON(i)
	}
}

func SendInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("send invoice " + id)
	}
}

func VoidInvoice(r *persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		switch i, err := r.FindInvoice(id); err {
		case nil:
			i.Status = invoice.StatusVoid
			i.VoidedAt = time.Now().Unix()
			return c.JSON(i)
		case persistence.ErrorInvoiceNotFound:
			return c.SendStatus(fasthttp.StatusNotFound)
		default:
			return err
		}
	}
}
