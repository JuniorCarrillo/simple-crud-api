package api

import (
	"github.com/JuniorCarrillo/simple-crud-api/api/handlers"
	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheck)
	e.GET("/", handlers.HealthCheck)

}

func Api(g *echo.Group) {
	g.POST("/team", handlers.TeamSearch)
	g.POST("/character", handlers.CharacterSave)
	g.GET("/character", handlers.CharacterSearch)
	g.GET("/character/:id", handlers.CharacterById)
	g.PUT("/character/:id", handlers.CharacterUpdateById)
	g.DELETE("/character/:id", handlers.CharacterDeleteById)
}
