package server

import "github.com/gofiber/fiber/v2"

func GetWelcome(c *fiber.Ctx) error {
	return c.Render("dashboard/welcome", fiber.Map{})
}
