package router

import (
	"github.com/JuniorCarrillo/simple-crud-api/api"
	"github.com/JuniorCarrillo/simple-crud-api/api/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	apiGroup := e.Group("/api/v1")

	middlewares.SetMainMiddleWares(e)
	middlewares.SetXApiKeyMiddlewares(apiGroup)

	api.MainGroup(e)
	api.Api(apiGroup)

	return e
}
