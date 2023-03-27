package tools

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strings"
)

func ToUrl(ctx echo.Context) string {
	uri := FilterTeam(ctx.Request().RequestURI)
	host := ctx.Request().Host
	protocol := ctx.Request().TLS

	url := fmt.Sprintf("http://%s%s", host, TotalUri(uri))
	if protocol != nil {
		url = fmt.Sprintf("https://%s%s", host, TotalUri(uri))
	}

	return url
}

func TotalUri(str string) string {
	uri := strings.Split(str, "?")
	if len(uri) > 1 {
		return uri[0]
	}
	return str
}

func FilterTeam(str string) string {
	return strings.Replace(str, "team", "character", -1)
}
