package controllers

import (
	"net/http"

	"echo_rest/models"

	"github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
	result, err := models.HealthCheck()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
