package handlers

import (
	"github.com/JuniorCarrillo/simple-crud-api/api/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	msg := "The service is alive!"
	response := dtos.Response{
		Message: &msg,
	}

	return c.JSON(http.StatusOK, response)
}
