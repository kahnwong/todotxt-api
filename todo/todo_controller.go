package todo

import (
	"github.com/gofiber/fiber/v2"
)

func TodoTodayController(c *fiber.Ctx) error {
	return c.JSON(getTodos())
}
