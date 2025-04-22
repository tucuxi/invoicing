package handlers

import (
	"errors"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tucuxi/invoicing/internal/pkg/persistence"
	"github.com/tucuxi/invoicing/pkg/invoice"
)

type PayInvoiceParameters struct {
	PaidOutOfBand bool  `json:"paid_out_of_band" form:"paid_out_of_band"`
	AmountPaid    int64 `json:"amount_paid" form:"amount_paid"`
}

func CreateInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		i := new(invoice.Invoice)

		if c.BodyParser(i) != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if slices.Index(invoice.InvoiceTypes, i.Type) == -1 {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if i.Recipient == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		i.ID = invoice.NewInvoiceID()
		i.Status = invoice.StatusDraft
		i.DraftedAt = time.Now().Unix()
		err := r.CreateInvoice(i)
		if err != nil {
			return err
		}
		return c.JSON(i)
	}
}

func UpdateInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := r.FindInvoice(id)
		if errors.Is(err, invoice.ErrorInvoiceNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		if err != nil {
			return err
		}
		if i.Status != invoice.StatusDraft {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		j := *i
		err = c.BodyParser(&j)
		if err != nil {
			return err
		}
		if j.ID != i.ID ||
			j.Type != i.Type ||
			j.DraftedAt != i.DraftedAt ||
			j.PaidAt != i.PaidAt ||
			j.VoidedAt != i.VoidedAt ||
			j.MarkedUncollectibleAt != i.MarkedUncollectibleAt {

			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = r.UpdateInvoice(&j)
		if err != nil {
			return err
		}
		return c.JSON(&j)
	}
}

func RetrieveUpcomingInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := c.FormValue("customer")
		return c.SendString("retrieve upcoming invoice for customer " + customer)
	}
}

func RetrieveInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := r.FindInvoice(id)
		if errors.Is(err, invoice.ErrorInvoiceNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		if err != nil {
			return err
		}
		return c.JSON(i)
	}
}

func UpdateLineItem(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		line := c.Params("line")
		return c.SendString("update line " + line + " of invoice " + id)
	}
}

func RetrieveLineItems(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("retrieve line items for invoice " + id)
	}
}

func DeleteDraftInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := r.DeleteDraftInvoice(id)
		switch {
		case errors.Is(err, invoice.ErrorInvoiceNotFound):
			return c.SendStatus(fiber.StatusNotFound)
		case errors.Is(err, invoice.ErrorDeletionNotAllowed):
			return c.SendStatus(fiber.StatusBadRequest)
		case err != nil:
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func FinalizeInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("finalize invoice " + id)
	}
}

func MarkInvoiceUncollectible(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := r.FindInvoice(id)
		if errors.Is(err, invoice.ErrorInvoiceNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		if err != nil {
			return err
		}
		i.Status = invoice.StatusUncollectible
		i.MarkedUncollectibleAt = time.Now().Unix()
		err = r.UpdateInvoice(i)
		if err != nil {
			return err
		}
		return c.JSON(i)
	}
}

func PayInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := new(PayInvoiceParameters)

		if c.BodyParser(p) != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if !p.PaidOutOfBand {
			return c.SendStatus(fiber.StatusNotImplemented)
		}

		id := c.Params("id")

		i, err := r.FindInvoice(id)
		if errors.Is(err, invoice.ErrorInvoiceNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		if err != nil {
			return err
		}
		i.Status = invoice.StatusPaid
		i.PaidAt = time.Now().Unix()
		i.PaidOutOfBand = p.PaidOutOfBand
		i.AmountPaid = p.AmountPaid
		err = r.UpdateInvoice(i)
		if err != nil {
			return err
		}
		return c.JSON(i)
	}
}

func SendInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("send invoice " + id)
	}
}

func VoidInvoice(r persistence.InvoiceRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := r.FindInvoice(id)
		if errors.Is(err, invoice.ErrorInvoiceNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		if err != nil {
			return err
		}
		i.Status = invoice.StatusVoid
		i.VoidedAt = time.Now().Unix()
		err = r.UpdateInvoice(i)
		if err != nil {
			return err
		}
		return c.JSON(i)
	}
}
