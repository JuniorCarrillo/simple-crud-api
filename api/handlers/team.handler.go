package handlers

import (
	"github.com/JuniorCarrillo/simple-crud-api/api/dtos"
	"github.com/JuniorCarrillo/simple-crud-api/api/tools"
	"github.com/JuniorCarrillo/simple-crud-api/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TeamSearch(ctx echo.Context) error {
	body := dtos.TeamRequest{}
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	if body.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}

	name := tools.ToCapitalize(body.Name)

	var limit int64 = 20
	if body.Limit != nil {
		limit = *body.Limit
	}

	page := body.Pages
	if body.Pages == 0 {
		page = 1
	}

	url := tools.ToUrl(ctx)
	response, err := database.Db().CharactersSearch(name, limit, page, url)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}
