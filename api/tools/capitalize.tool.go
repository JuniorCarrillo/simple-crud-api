package tools

import (
	"strings"
)

func ToCapitalize(str string) string {
	return strings.Title(strings.ToLower(str))
}
