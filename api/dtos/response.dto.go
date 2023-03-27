package dtos

import (
	"github.com/JuniorCarrillo/simple-crud-api/api/models"
)

type Response struct {
	Message   *string           `json:"message,omitempty"`
	Character *models.Character `json:"character,omitempty"`
}

type PaginationInfo struct {
	Count int64  `json:"count"`
	Pages int64  `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}
