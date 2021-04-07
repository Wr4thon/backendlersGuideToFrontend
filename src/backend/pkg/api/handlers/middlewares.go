package handlers

import (
	"net/http"

	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/services/animals"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (zoo *zooAPI) LoadAnimal() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			id, err := uuid.Parse(ctx.Param("animalID"))

			if err != nil {
				return ctx.NoContent(http.StatusBadRequest)
			}

			var animal animals.Animal
			if animal, err = zoo.animalsService.GetAnimal(id); err != nil {
				if errors.Is(errors.Cause(err), animals.ErrAnimalNotFound) {
					return ctx.NoContent(http.StatusNotFound)
				}

				return echo.NotFoundHandler(ctx)
			}

			ctx = Context{
				Context: ctx,
				Animal:  animal,
			}

			return next(ctx)
		}
	}
}
