package animals

import (
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/repository"
	"github.com/google/uuid"
)

type (
	Animal struct {
		ID         uuid.UUID
		Name       string
		Species    string
		Properties map[string]interface{}
	}

	AnimalUpdate struct {
		Properties map[string]interface{}
	}
)

func newServiceAnimal(animal repository.Animal) Animal {
	return Animal{
		ID:         animal.ID,
		Name:       animal.Name,
		Properties: animal.Properties,
		Species:    animal.Species,
	}
}

func newRepositoryAnimal(animal Animal) repository.Animal {
	return repository.Animal{
		ID:         animal.ID,
		Name:       animal.Name,
		Properties: animal.Properties,
		Species:    animal.Species,
	}
}
