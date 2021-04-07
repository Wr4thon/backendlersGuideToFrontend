package inmemory

import (
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/repository"
	"github.com/google/uuid"
)

type (
	animals struct {
		cache map[uuid.UUID]repository.Animal
	}
)

func New(loadDemoData bool) repository.Animals {
	cache := demoData(loadDemoData)

	return &animals{
		cache: cache,
	}
}

func (repo *animals) Get() ([]repository.Animal, error) {
	animals := make([]repository.Animal, len(repo.cache))

	var i int
	for _, animal := range repo.cache {
		animals[i] = animal
		i++
	}

	return animals, nil
}

func (repo *animals) GetByID(id uuid.UUID) (repository.Animal, error) {
	if animal, ok := repo.cache[id]; ok {
		return animal, nil
	}

	return repository.Animal{}, repository.ErrAnimalNotFound
}

func (repo *animals) Upsert(animal repository.Animal) (uuid.UUID, error) {
	if animal.ID == uuid.Nil {
		animal.ID = uuid.New()
	}

	repo.cache[animal.ID] = animal
	return animal.ID, nil
}

func (repo *animals) Delete(id uuid.UUID) error {
	delete(repo.cache, id)
	return nil
}
