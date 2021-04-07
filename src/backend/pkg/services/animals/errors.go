package animals

import "github.com/pkg/errors"

var (
	ErrAnimalNotFound error = errors.New("animal not found")
)
