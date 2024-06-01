package routes

import (
	"github.com/harish876/portfolio/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterCommandRoutes(e *echo.Echo) {
	apiGroup := e.Group("/commands")
	apiGroup.POST("", handlers.CommandHandler)
}
