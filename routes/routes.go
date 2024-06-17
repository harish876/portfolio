package routes

import (
	"github.com/harish876/portfolio/app"
	"github.com/harish876/portfolio/app/about"
	api_projects "github.com/harish876/portfolio/app/api/v1/projects"
	"github.com/harish876/portfolio/app/commands"
	"github.com/harish876/portfolio/app/project"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup0 := e.Group("/about")
	apiGroup0.GET("", about.AboutHandler)

	apiGroup1 := e.Group("/api/v1/projects")
	apiGroup1.GET("", api_projects.GetProjectsDataFromGithub)

	apiGroup2 := e.Group("/commands")
	apiGroup2.POST("", commands.CommandHandler)

	apiGroup3 := e.Group("")
	apiGroup3.GET("", app.HomeHandler)

	apiGroup4 := e.Group("/project")
	apiGroup4.GET("", project.ProjectHandler)
}
