package main

import (
	"github.com/gofiber/fiber"
	// jwtware
)

func createNew(c *fiber.Ctx) {
	if c.Body() == "" {
		c.JSON(fiber.Map{
			"message": "You need to have a body in the request",
		})
	} else {
		c.JSON(fiber.Map{
			"message": "Hello There new user",
		})
	}
}
func list(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"Name": "User 1",
	})
}
