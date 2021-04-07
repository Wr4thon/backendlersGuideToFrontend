package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/services/animals"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type (
	ZooAPI interface {
		LoadAnimal() echo.MiddlewareFunc

		GetAnimals(ctx echo.Context) error
		AddAnimal(ctx echo.Context) error

		GetAnimal(ctx echo.Context) error
		UpdateAnimal(ctx echo.Context) error
		DeleteAnimal(ctx echo.Context) error
	}

	zooAPI struct {
		animalsService animals.Service
	}
)

func NewZooAPI(animalsService animals.Service) ZooAPI {
	return &zooAPI{
		animalsService: animalsService,
	}
}

func (zoo *zooAPI) GetAnimals(ctx echo.Context) error {
	allAnimals, err := zoo.animalsService.GetAnimals()
	if err != nil {
		if errors.Is(errors.Cause(err), animals.ErrAnimalNotFound) {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.NoContent(http.StatusInternalServerError)
	}

	if len(allAnimals) == 0 {
		return ctx.NoContent(http.StatusNoContent)
	}

	return ctx.JSON(http.StatusOK, allAnimals)
}

func (zoo *zooAPI) AddAnimal(ctx echo.Context) error {
	request := AddAnimalRequest{}
	if err := readBody(ctx.Request().Body, &request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	animal := Animal{
		Name:       request.Name,
		Species:    request.Species,
		Properties: request.Properties,
	}

	var err error
	animal.ID, err = zoo.animalsService.AddAnimal(newServiceAnimal(animal))

	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusCreated, animal)
}

func (zoo *zooAPI) GetAnimal(ctx echo.Context) error {
	c := ctx.(Context)
	return ctx.JSON(http.StatusOK, newAnimal(c.Animal))
}

func (zoo *zooAPI) UpdateAnimal(ctx echo.Context) error {
	c := ctx.(Context)

	request := UpdateAnimalRequest{}
	if err := readBody(ctx.Request().Body, &request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	a := c.Animal
	update := animals.AnimalUpdate{
		Properties: request.Properties,
	}

	if err := zoo.animalsService.UpdateAnimal(a, update); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusAccepted)
}

func (zoo *zooAPI) DeleteAnimal(ctx echo.Context) error {
	c := ctx.(Context)
	if err := zoo.animalsService.DeleteAnimal(c.Animal.ID); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusAccepted)
}

func readBody(reader io.Reader, result interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, result); err != nil {
		return err
	}
	return nil
}
