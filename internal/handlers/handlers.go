package handlers

import "github.com/gofiber/fiber/v2"

func CreateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("create invoice")
	}
}

func UpdateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("update invoice " + id)
	}
}

func RetrieveUpcomingInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := c.FormValue("customer")
		return c.SendString("retrieve upcoming invoice for customer " + customer)
	}
}

func RetrieveInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("retrieve invoice " + id)
	}
}

func UpdateLineItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		line := c.Params("line")
		return c.SendString("update line " + line + " of invoice " + id)
	}
}

func RetrieveLineItems() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("retrieve line items for invoice " + id)
	}
}

func DeleteDraftInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("delete draft invoice " + id)
	}
}

func FinalizeInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("finalize invoice " + id)
	}
}

func MarkInvoiceUncollectible() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("mark invoice " + id + " uncollectible")
	}
}

func PayInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("pay invoice " + id)
	}
}

func SendInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("send invoice " + id)
	}
}

func VoidInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("void invoice " + id)
	}
}
