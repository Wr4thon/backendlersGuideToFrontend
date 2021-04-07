package main

import (
	"log"
	"net/http"

	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/api"
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/api/handlers"
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/repository/inmemory"
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/services/animals"
	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/backendlersGuideToFrontend-presentation/src/backend/docs"
	"github.com/labstack/echo/v4/middleware"
)

// @title Zoo API documentation
// @version 0.0
// @description This is the documentation for the zoo service v0
// @contact.email me@wr4thon.de
// @BasePath /api/v0
//
func main() {
	e := echo.New()

	e.Use(
		middleware.RemoveTrailingSlash(),
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/docs", func(c echo.Context) error {
		http.Redirect(c.Response(), c.Request(), "/docs/index.html", http.StatusMovedPermanently)
		return nil
	})

	animalsRepo := inmemory.New(true)
	animalsService := animals.New(animalsRepo)
	zoo := handlers.NewZooAPI(animalsService)

	api.RegisterRoutes(e, zoo)
	log.Panic(e.Start(":8080"))
}
