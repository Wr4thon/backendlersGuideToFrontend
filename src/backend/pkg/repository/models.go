package repository

import (
	"github.com/google/uuid"
)

type (
	Animal struct {
		ID         uuid.UUID              `json:"id" bson:"_id"`
		Name       string                 `json:"name,omitempty" bson:"name,omitempty"`
		Species    string                 `json:"species,omitempty" bson:"species,omitempty"`
		Properties map[string]interface{} `json:"properties,omitempty" bson:"properties,omitempty"`
	}
)
