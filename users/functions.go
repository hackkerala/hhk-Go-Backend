package users

import (
	"log"
	"net/http"
	"time"

	"github.com/athul/anonblog/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Newuser(c *gin.Context) {
	dbc := db.DBConn
	var input = new(db.User)
	//c.BindJSON(input)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//nuser := db.User{Name: input.Name, Email: input.Email}
	if input.Emailval() == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a valid email address"})
		return
	}
	dbc.Create(&input)
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	//log.Println(ll.Value)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = input.Name
	log.Println(input.Name)
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": t, "user": input})
}
func GetUsers(c *gin.Context) {
	dbc := db.DBConn

	defer dbc.Close()

	var users []db.User
	dbc.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func FindUserbyID(c *gin.Context) { // Get model if exist
	var user db.User
	db := db.DBConn
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
