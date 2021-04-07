package animals

import (
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	Service interface {
		GetAnimals() ([]Animal, error)
		GetAnimal(uuid.UUID) (Animal, error)
		UpdateAnimal(Animal, AnimalUpdate) error
		DeleteAnimal(uuid.UUID) error
		AddAnimal(Animal) (uuid.UUID, error)
	}

	service struct {
		animalsRepo repository.Animals
	}
)

func New(animalsRepo repository.Animals) Service {
	return &service{
		animalsRepo: animalsRepo,
	}
}

func (service *service) GetAnimals() ([]Animal, error) {
	animals, err := service.animalsRepo.Get()

	if err != nil {
		if errors.Is(errors.Cause(err), repository.ErrAnimalNotFound) {
			return nil, ErrAnimalNotFound
		}

		return nil, errors.Wrap(err, "error while getting animals from repository")
	}

	result := make([]Animal, len(animals))
	for i := range animals {
		result[i] = newServiceAnimal(animals[i])
	}

	return result, nil
}

func (service *service) GetAnimal(id uuid.UUID) (Animal, error) {
	animal, err := service.animalsRepo.GetByID(id)

	if err != nil {
		if errors.Is(errors.Cause(err), repository.ErrAnimalNotFound) {
			return Animal{},
				errors.Wrapf(
					ErrAnimalNotFound,
					"animal with id %s not found in repository",
					id.String(),
				)
		}

		return Animal{},
			errors.Wrapf(
				err,
				"error while getting animal with id %s from repository",
				id.String(),
			)
	}

	return newServiceAnimal(animal), nil
}

func (service *service) UpdateAnimal(animal Animal, update AnimalUpdate) error {
	for prop, val := range update.Properties {
		animal.Properties[prop] = val
	}

	_, err := service.animalsRepo.Upsert(newRepositoryAnimal(animal))
	return err
}

func (service *service) DeleteAnimal(id uuid.UUID) error {
	return service.animalsRepo.Delete(id)
}

func (service *service) AddAnimal(animal Animal) (uuid.UUID, error) {
	return service.animalsRepo.Upsert(newRepositoryAnimal(animal))
}
