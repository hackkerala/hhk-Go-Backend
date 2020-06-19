package api

import (
	"net/http"

	"github.com/athul/anonblog/db"
	"github.com/gin-gonic/gin"
)

func Newhack(c *gin.Context){
	dbc:=db.DBConn
	var h = new(db.Hackathon)

	if err:= c.ShouldBindJSON(h); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error()	
		})
		return
	}
	if h.Emaileval() != false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a valid email address"})
		return
	}
	dbc.Create(&h)
	c.JSON(http.StatusCreated, gin.H{
		"data":h
	})
}