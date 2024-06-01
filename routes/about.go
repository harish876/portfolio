package routes

import (
	"github.com/harish876/portfolio/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAboutRoutes(e *echo.Echo) {
	apiGroup := e.Group("/about")

	apiGroup.GET("", handlers.AboutHandler)
	apiGroup.POST("", handlers.AboutHandler)
}
