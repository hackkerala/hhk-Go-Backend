package db

import (
	"fmt"
	"os/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DBConn *gorm.DB

func ConnectDB() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database Migrated")
}
