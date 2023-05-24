package controllers

import (
	"net/http"

	"echo_rest/models"

	"github.com/labstack/echo"
)

func FetchAllEmployee(c echo.Context) error {
	result, err := models.FetchAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostEmployee(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone_number := c.FormValue("phone_number")

	result, err := models.PostEmployee(name,address,phone_number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PutEmployee(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone_number := c.FormValue("phone_number")

	result, err := models.PutEmployee(id,name,address,phone_number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DelEmployee(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DelEmployee(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
