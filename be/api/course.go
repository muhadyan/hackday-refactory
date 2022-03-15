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

// Post Data Course
func PostCourse(c echo.Context) error {
	db := db.DbManager()
	// Validate Input
	var input model.Courses
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if input.CourseName == "" {
		return c.JSON(http.StatusNoContent, "Data not complete or empty")
	}

	course := model.Courses{CourseName: input.CourseName}

	if err := db.Create(&course).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}

// Update Data Course
func UpdateCourse(c echo.Context) error {
	db := db.DbManager()
	// Get model if exist
	var course model.Courses
	param := c.Param("course_id")

	if err := db.Where("course_id = ?", param).First(&course).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Record not found!")
	}

	// validate input
	var input model.Courses
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Model(&course).Where("course_id = ?", param).Updates(input)

	return c.JSON(http.StatusOK, course)
}

// Delete Data Students
func DeleteCourse(c echo.Context) error {
	db := db.DbManager()
	var course model.Courses
	param := c.Param("course_id")

	if err := db.Where("course_id = ?", param).First(&course).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Record not found")
	}

	db.Where("course_id = ?", param).Delete(&course)

	return c.JSON(http.StatusOK, "Data has been deleted")
}