package users

import (
	"fmt"
	"time"

	"github.com/athul/anonblog/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

//User type for anonymous user
type User struct {
	gorm.Model

	Name string `json:"username"`
}

// Article type
type Article struct {
	gorm.Model

	Author User   `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Newuser(c *fiber.Ctx) {
	db := db.DBConn
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&user)
	fmt.Println(c.BodyParser(user))
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(token)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"token": t, "user": user})
}
