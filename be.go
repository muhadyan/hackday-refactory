package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	connect()

	r.GET("/students", getStudents)

	r.Run()
}

type Students struct {
	StudentID int64  `json:"StudentID"`
	FullName  string `json:"FullName"`
	ExtraID   int64  `json:"ExtraID"`
}

var DB *gorm.DB

func connect() {
	dsn := "root:test@tcp(127.0.0.1:3306)/hackday"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}

// Get Data Students
func getStudents(c *gin.Context) {
	var students []Students
	DB.Find(&students)

	c.JSON(200, students)
}
