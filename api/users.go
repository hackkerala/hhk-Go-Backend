package api

import (
	"net/http"

	"github.com/athul/anonblog/db"
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
	c.JSON(http.StatusCreated, gin.H{"id": input.ID, "name": input.Name, "email": input.Email})
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
