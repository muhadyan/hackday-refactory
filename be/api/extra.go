package api

import (
	"be/db"
	"be/model"
	"net/http"

	"github.com/labstack/echo"
)

// Get Data Extras
func GetExtras(c echo.Context) error {
	db := db.DbManager()
	var extras []model.Extras
	db.Find(&extras)

	return c.JSON(http.StatusOK, extras)
}

// Post Data Extra
func PostExtra(c echo.Context) error {
	db := db.DbManager()
	// Validate Input
	var input model.Extras
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if input.ExtraName == "" {
		return c.JSON(http.StatusNoContent, "Data not complete or empty")
	}

	extra := model.Extras{ExtraName: input.ExtraName}

	if err := db.Create(&extra).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}

// Update Data Extra
func UpdateExtra(c echo.Context) error {
	db := db.DbManager()
	// Get model if exist
	var student model.Extras
	param := c.Param("extra_id")

	if err := db.Where("extra_id = ?", param).First(&student).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Record not found!")
	}

	// validate input
	var input model.Extras
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Model(&student).Where("extra_id = ?", param).Updates(input)

	return c.JSON(http.StatusOK, student)
}

// Delete Data Extra
func DeleteExtra(c echo.Context) error {
	db := db.DbManager()
	var extra model.Extras
	param := c.Param("extra_id")

	if err := db.Where("extra_id = ?", param).First(&extra).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Record not found")
	}

	db.Where("extra_id = ?", param).Delete(&extra)

	return c.JSON(http.StatusOK, "Data has been deleted")
}