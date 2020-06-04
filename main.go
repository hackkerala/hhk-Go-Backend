package main

import (
	"fmt"
	"os/user"

	"github.com/athul/anonblog/db"
	"github.com/athul/anonblog/users"
	"github.com/gofiber/fiber" // jwtware
	"github.com/jinzhu/gorm"
)

func initDatabase() {
	var err error
	db.DBConn, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	db.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database Migrated")
}
func main() {
	app := fiber.New()

	initDatabase()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World")
	})
	setupRoutes(app)
	app.Listen(3000)
}
func setupRoutes(app *fiber.App) {
	api := app.Group("/api", func(c *fiber.Ctx) {
		c.Set("X-Custom-Header", "isafgiwegfuiqwgfivfberik")
		c.Next()
	})
	// api.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	api.Post("/new", createNew)
	api.Get("/list", list)
	api.Post("/login", users.Newuser)

}
