package dtos

import (
	"fmt"
	"github.com/JuniorCarrillo/simple-crud-api/api/models"
	"github.com/JuniorCarrillo/simple-crud-api/api/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CharacterToSave(in *Character) *models.NewCharacter {
	now := time.Now()
	return &models.NewCharacter{
		Name:      in.Name,
		Status:    in.Status,
		Species:   in.Species,
		Type:      in.Type,
		Gender:    in.Gender,
		CreatedAt: now,
	}
}

func CharacterToUpdate(in *CharacterUpdate) *models.UpdateCharacter {
	now := time.Now()
	return &models.UpdateCharacter{
		Name:      in.Name,
		Status:    in.Status,
		Species:   in.Species,
		Type:      in.Type,
		Gender:    in.Gender,
		UpdatedAt: now,
	}
}

func CharacterToResponse(in *Character, out *mongo.InsertOneResult, now time.Time) *models.Character {
	return &models.Character{
		ID:        out.InsertedID.(primitive.ObjectID).Hex(),
		Name:      in.Name,
		Status:    in.Status,
		Species:   in.Species,
		Type:      in.Type,
		Gender:    in.Gender,
		CreatedAt: now,
		UpdatedAt: &now,
	}
}

func CharacterToPagination(characters []*models.Character, info, page, limit int64, url, search string) *CharacterPaginationResponse {
	pages := tools.ToPages(info, limit)
	next := fmt.Sprintf("%s?pages=%v", url, page+1)
	prev := ""

	if page > 1 {
		next = fmt.Sprintf("%s?pages=%v", url, page+1)
	}

	if page > pages {
		next = ""
	}

	if page-1 > 0 {
		prev = fmt.Sprintf("%s?pages=%v", url, page-1)
	}

	if search != "" {
		if page <= pages {
			next = fmt.Sprintf("%s&search=%s", next, search)
		}
		if page-1 > 0 {
			prev = fmt.Sprintf("%s&search=%s", prev, search)
		}
	}

	return &CharacterPaginationResponse{
		Result: characters,
		Info: PaginationInfo{
			Count: info,
			Pages: pages,
			Next:  next,
			Prev:  prev,
		},
	}
}

type Character struct {
	Name    string  `json:"name"`
	Status  *string `json:"status,omitempty"`
	Species *string `json:"species,omitempty"`
	Type    *string `json:"type,omitempty"`
	Gender  *string `json:"gender,omitempty"`
}

type CharacterUpdate struct {
	Name    *string `json:"name,omitempty"`
	Status  *string `json:"status,omitempty"`
	Species *string `json:"species,omitempty"`
	Type    *string `json:"type,omitempty"`
	Gender  *string `json:"gender,omitempty"`
}

type CharacterPaginationResponse struct {
	Info   PaginationInfo      `json:"info"`
	Result []*models.Character `json:"result,omitempty"`
}
