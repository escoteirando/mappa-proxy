package handlers

import "github.com/gofiber/fiber/v2"

type ReplyMessage struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func reply_status(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ReplyMessage{
		Message: message,
	})
}

func reply_error(c *fiber.Ctx, statusCode int, message string, err error) error {
	return c.Status(statusCode).JSON(ReplyMessage{
		Message: message,
		Error:   err.Error(),
	})
}
