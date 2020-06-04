package main

import (
	"github.com/dgrijalva/jwt-go"

  "github.com/gofiber/fiber"
  "github.com/gofiber/jwt" // jwtware
)

func api(app *fiber.App) {

	api := app.Group("/api", func(c *fiber.Ctx) {
		c.Set("X-Custom-Header", "isafgiwegfuiqwgfivfberik")
		c.Next()
	})
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	  }))
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
	func login(c *fiber.Ctx){
		user:=c.Formvak
	}
}
