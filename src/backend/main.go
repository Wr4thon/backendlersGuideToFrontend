package main

import (
	"log"

	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/api"
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/api/handlers"
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/repository/inmemory"
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/services/animals"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(
		middleware.RemoveTrailingSlash(),
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	animalsRepo := inmemory.New(true)
	animalsService := animals.New(animalsRepo)
	zoo := handlers.NewZooAPI(animalsService)

	api.RegisterRoutes(e, zoo)
	log.Panic(e.Start(":8080"))
}
