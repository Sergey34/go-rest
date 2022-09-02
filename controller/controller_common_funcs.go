package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func respondWithJson(c *fiber.Ctx, code int, payload interface{}) error {
	if payload == nil {
		return c.Status(code).Send(nil)
	}
	return c.Status(code).JSON(payload)
}

func respondWithError(c *fiber.Ctx, code int, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return respondWithJson(c, code, map[string]string{"error": msg})
}
