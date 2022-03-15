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