package handlers

import (
	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/services/animals"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Animal animals.Animal
}
