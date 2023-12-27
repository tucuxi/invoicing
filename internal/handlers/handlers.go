package handlers

import "github.com/gofiber/fiber/v2"

func CreateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("create invoice")
	}
}

func UpdateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		invoice := c.Params("invoice")
		return c.SendString("update invoice " + invoice)
	}
}

func UpdateLineItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		invoice := c.Params("invoice")
		line := c.Params("line")
		return c.SendString("update line " + line + " of invoice " + invoice)
	}
}

func RetrieveInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		invoice := c.Params("invoice")
		return c.SendString("retrieve invoice " + invoice)
	}
}
