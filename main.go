package main

import (
	"github.com/gofiber/fiber"
	"github.com/athul/anonblog/db"
)
func initDatabase() {
	var err error
	db.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
func main() {
	app := fiber.New()
	
	initDatabase()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World")
	})
	api(app)
	app.Listen(3000)
}
