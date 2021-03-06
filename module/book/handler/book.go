package handler

import "github.com/gofiber/fiber/v2"

func (b *HttpBookHandler) BookList(c *fiber.Ctx) error {
	return c.SendString("book list")
}
