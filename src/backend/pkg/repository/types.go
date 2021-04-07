package repository

import "github.com/google/uuid"

type (
	Animals interface {
		Get() ([]Animal, error)
		Upsert(Animal) (uuid.UUID, error)
		GetByID(uuid.UUID) (Animal, error)
		Delete(uuid.UUID) error
	}
)
