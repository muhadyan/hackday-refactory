package api

import (
	"be/db"
	"be/model"
	"be/model/request"
	"net/http"

	"github.com/labstack/echo"
)

// Get Data Students
func GetStudents(c echo.Context) error {
	db := db.DbManager()
	var students []model.Students
	db.Table(`students`).
		Select(`students.student_id, students.full_name, extras.extra_name`).
		Joins("inner join extras on extras.extra_id = students.extra_id").
		Scan(&students)

	return c.JSON(http.StatusOK, students)
}

// Post Data Students
func PostStudent(c echo.Context) error {
	db := db.DbManager()
	// Validate Input
	var input request.Students
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	student := request.Students{
		StudentID: input.StudentID,
		FullName:  input.FullName,
		ExtraID:   input.ExtraID}

	db.Create(&student)

	return c.JSON(http.StatusOK, "Success")
}

// Update Data Students
func UpdateStudent(c echo.Context) error {
	db := db.DbManager()
	// Get model if exist
	var student model.Students
	param := c.Param("student_id")

	if err := db.Where("student_id = ?", param).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, "Record not found!")
	}

	// validate input
	var input model.Students
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Model(&student).Updates(input)

	return c.JSON(http.StatusOK, student)
}

// Delete Data Students
func DeleteStudent(c echo.Context) error {
	db := db.DbManager()
	var student model.Students
	param := c.Param("student_id")

	if err := db.Where("student_id = ?", param).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, "Record not found")
	}

	db.Delete(&student)

	return c.JSON(http.StatusOK, "Data has been deleted")
}
