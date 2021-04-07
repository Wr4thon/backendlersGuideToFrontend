package handlers

import (
	service "github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/services/animals"
	"github.com/google/uuid"
)

type (
	Animal struct {
		ID         uuid.UUID              `json:"id"`
		Name       string                 `json:"name,omitempty"`
		Species    string                 `json:"species,omitempty"`
		Properties map[string]interface{} `json:"properties,omitempty"`
	}
)

func newAnimal(animal service.Animal) Animal {
	return Animal{
		ID:         animal.ID,
		Name:       animal.Name,
		Species:    animal.Name,
		Properties: animal.Properties,
	}
}

func newServiceAnimal(animal Animal) service.Animal {
	return service.Animal{
		ID:         animal.ID,
		Name:       animal.Name,
		Species:    animal.Name,
		Properties: animal.Properties,
	}
}
