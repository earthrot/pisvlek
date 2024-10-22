package server

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	users, _ := q.GetUsers(ctx)

	return c.Render("users/userlist", fiber.Map{
		"Users": users,
	})
}
