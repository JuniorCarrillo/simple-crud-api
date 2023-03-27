package models

import "time"

type NewCharacter struct {
	Name      string    `json:"name" bson:"name"`
	Status    *string   `json:"status,omitempty" bson:"status,omitempty"`
	Species   *string   `json:"species,omitempty" bson:"species,omitempty"`
	Type      *string   `json:"type,omitempty" bson:"type,omitempty"`
	Gender    *string   `json:"gender,omitempty" bson:"gender,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type UpdateCharacter struct {
	Name      *string   `json:"name,omitempty" bson:"name,omitempty"`
	Status    *string   `json:"status,omitempty" bson:"status,omitempty"`
	Species   *string   `json:"species,omitempty" bson:"species,omitempty"`
	Type      *string   `json:"type,omitempty" bson:"type,omitempty"`
	Gender    *string   `json:"gender,omitempty" bson:"gender,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Character struct {
	ID        string     `json:"_id" bson:"_id"`
	Name      string     `json:"name" bson:"name"`
	Status    *string    `json:"status,omitempty" bson:"status,omitempty"`
	Species   *string    `json:"species,omitempty" bson:"species,omitempty"`
	Type      *string    `json:"type,omitempty" bson:"type,omitempty"`
	Gender    *string    `json:"gender,omitempty" bson:"gender,omitempty"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
