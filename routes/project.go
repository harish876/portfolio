package routes

import (
	"github.com/harish876/portfolio/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterProjectRoutes(e *echo.Echo) {
	apiGroup := e.Group("/project")

	apiGroup.GET("", handlers.ProjectHandler)
}

func RegisterProjectApiRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api/v1")

	apiGroup.GET("/projects", handlers.GetProjectsDataFromGithub)
}
