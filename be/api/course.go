package api

import (
	"be/db"
	"be/model"
	"net/http"

	"github.com/labstack/echo"
)

// Get Data Courses
func GetCourses(c echo.Context) error {
	db := db.DbManager()
	var courses []model.Courses
	db.Find(&courses)

	return c.JSON(http.StatusOK, courses)
}