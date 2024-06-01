package routes

import (
	"github.com/harish876/portfolio/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterHomeRoutes(e *echo.Echo) {
	apiGroup := e.Group("/")

	apiGroup.GET("", handlers.HomeHandler)
}
