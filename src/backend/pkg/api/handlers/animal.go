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

// @tags animals
// @Router /animals [GET]
// @Summary get all animals
// @Description get all animals, that are contained in the database
// @Produce json
// @Success 200 {array} Animal "when the database read was successfull"
// @Success 204 {nil} nil "when no animals are in the database."
// @Failure 500 {nil} nil "when the database returns an error"
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

	animals := make([]Animal, len(allAnimals))

	for i := range allAnimals {
		animals[i] = newAnimal(allAnimals[i])
	}

	return ctx.JSON(http.StatusOK, animals)
}

// @tags animals
// @Router /animals [POST]
// @Summary add a new animal
// @Description add a new animal to the database
// @Produce json
// @Consumes json
// @Header 201 {string} Location "/api/v0/animals/{animalID}"
// @Success 201 {object} Animal "the animal was created and can be found under the url in the Location header"
// @Success 400 {nil} nil "there was an error with the request; Check the json you sent"
// @Failure 500 {nil} nil "when the database returns an error"
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

	ctx.Response().Header().Add("Location", "/api/v0/animals/"+animal.ID.String())

	return ctx.JSON(http.StatusCreated, animal)
}

// @tags animals
// @Router /animals/{animalID} [GET]
// @Summary get a specific animal
// @Param animalID path string true "the uuid of the animal in the database"
// @Description get a animal specified by the {animalID} parameter in the URL
// @Produce json
// @Success 200 {object} Animal "when the database read was successfull"
// @Success 400 {nil} nil "when the {animalID} is not a valid uuid.UUID"
// @Success 404 {nil} nil "no animal was found with the provided {animalID}"
// @Failure 500 {nil} nil "when the database returns an error"
func (zoo *zooAPI) GetAnimal(ctx echo.Context) error {
	c := ctx.(Context)
	return ctx.JSON(http.StatusOK, newAnimal(c.Animal))
}

// @tags animals
// @Router /animals/{animalID} [PUT]
// @Summary update a specific animal
// @Param animalID path string true "the uuid of the animal in the database"
// @Param document body UpdateAnimalRequest true "the data you want to update for the animal"
// @Description update an animal specified by the {animalID} parameter in the URL
// @Consumes json
// @Success 200 {object} Animal "when the database read was successfull"
// @Success 400 {nil} nil "when the {animalID} is not a valid uuid.UUID"
// @Success 400 {nil} nil "the request was not valid"
// @Success 404 {nil} nil "no animal was found with the provided {animalID}"
// @Failure 500 {nil} nil "when the database returns an error"
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

// @tags animals
// @Router /animals/{animalID} [DELETE]
// @Summary delete a specific animal
// @Param animalID path string true "the uuid of the animal in the database"
// @Description delete an animal specified by the {animalID} parameter in the URL
// @Success 202 {nil} nil "when the operation was successfull"
// @Success 400 {nil} nil "when the {animalID} is not a valid uuid.UUID"
// @Success 404 {nil} nil "no animal was found with the provided {animalID}"
// @Failure 500 {nil} nil "when the database returns an error"
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
