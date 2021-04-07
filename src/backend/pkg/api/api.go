package api

import (
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/api/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, zoo handlers.ZooAPI) {
	api := e.Group("/api/v0")

	animals := api.Group("/animals")

	loadAnimal := zoo.LoadAnimal()

	animals.GET("", zoo.GetAnimals)
	animals.POST("", zoo.AddAnimal)
	animals.GET("/:animalID", zoo.GetAnimal, loadAnimal)
	animals.DELETE("/:animalID", zoo.DeleteAnimal, loadAnimal)
	animals.PUT("/:animalID", zoo.UpdateAnimal, loadAnimal)
}
