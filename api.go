package main

import "github.com/gofiber/fiber"

func api(app *fiber.App) {

	api := app.Group("/api", func(c *fiber.Ctx) {
		c.Set("X-Custom-Header", "isafgiwegfuiqwgfivfberik")
		c.Next()
	})
	api.Post("/new", func(c *fiber.Ctx) {
		if c.Body() == "" {
			c.JSON(fiber.Map{
				"message": "You need to have a body in the request",
			})
		} else {
			c.JSON(fiber.Map{
				"message": "Hello There new user",
			})
		}
	})
	api.Get("/list", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"Name": "User 1",
		})
	})
}
