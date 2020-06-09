package main

import (
	"net/http"

	"github.com/athul/anonblog/db"
	"github.com/athul/anonblog/users"
	"github.com/gin-gonic/gin" // jwtware
)

func main() {
	// app := fiber.New()
	r := gin.Default()
	db.ConnectDB()

	setupRoutes(r)
	// app.Listen(3000)

	r.Run()
}
func setupRoutes(r *gin.Engine) {
	// api := app.Group("/api", func(c *fiber.Ctx) {
	// 	c.Set("X-Custom-Header", "isafgiwegfuiqwgfivfberik")
	// 	c.Next()
	// })

	// api.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	// r.POST("/new", createNew)
	// r.GET("/list", list)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.POST("/login", users.Newuser)
	r.GET("/users.all", users.GetUsers)
	r.GET("/users/:id", users.FindBook)

}
