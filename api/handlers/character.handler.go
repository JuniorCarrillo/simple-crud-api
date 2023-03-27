package handlers

import (
	"errors"
	"fmt"
	"github.com/JuniorCarrillo/simple-crud-api/api/dtos"
	"github.com/JuniorCarrillo/simple-crud-api/api/tools"
	"github.com/JuniorCarrillo/simple-crud-api/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CharacterSearch(ctx echo.Context) error {
	pageQuery := ctx.QueryParam("pages")
	limitQuery := ctx.QueryParam("limit")
	searchQuery := ctx.QueryParam("search")

	var page int64 = 1
	if pageQuery != "" {
		pageValid, err := strconv.ParseInt(pageQuery, 10, 64)
		if err != nil {
			msg := fmt.Sprintf("The 'pages' parameter %d of type %T, must be positive numeric greater than 1", page, page)
			return echo.NewHTTPError(http.StatusBadRequest, msg)
		}
		page = pageValid
	}

	var limit int64 = 10
	if limitQuery != "" {
		limitValid, err := strconv.ParseInt(pageQuery, 10, 64)
		if err != nil {
			msg := fmt.Sprintf("The 'limit' parameter %d of type %T, must be positive numeric greater than 1", limit, limit)
			return echo.NewHTTPError(http.StatusBadRequest, msg)
		}
		limit = limitValid
	}

	url := tools.ToUrl(ctx)
	search := tools.ToCapitalize(searchQuery)
	response, err := database.Db().CharactersSearch(search, limit, page, url)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}

func CharacterSave(ctx echo.Context) error {
	body := dtos.Character{}
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	if body.Name == "" {
		return echo.NewHTTPError(http.StatusConflict, "Name is required")
	}

	response, err := database.Db().SaveCharacter(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := fmt.Sprintf("Item saved with id '%s'.", response.ID)
	return ctx.JSON(http.StatusCreated, dtos.Response{
		Message:   &msg,
		Character: response,
	})
}

func CharacterById(ctx echo.Context) error {
	id := ctx.Param("id")
	response, err := database.Db().CharacterById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := fmt.Sprintf("Item found for id '%s'.", response.ID)
	return ctx.JSON(http.StatusOK, dtos.Response{
		Message:   &msg,
		Character: response,
	})
}

func CharacterUpdateById(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("you need to provide an update id"))
	}

	body := dtos.CharacterUpdate{}
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	response, err := database.Db().CharacterUpdateById(id, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := fmt.Sprintf("The element with the id '%s' has been updated successfully.", response.ID)
	return ctx.JSON(http.StatusOK, dtos.Response{
		Message: &msg,
	})
}

func CharacterDeleteById(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("you need to provide an delete id"))
	}

	response, err := database.Db().CharacterDeleteById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := fmt.Sprintf("The element with the id '%s' has been deleted successfully.", response.ID)
	return ctx.JSON(http.StatusOK, dtos.Response{
		Message: &msg,
	})
}
