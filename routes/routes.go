package routes

import (
	api_projects "github.com/harish876/portfolio/app/api/v1/projects"
	"github.com/harish876/portfolio/app/commands"
	"github.com/harish876/portfolio/app/project"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup0 := e.Group("/about")
	apiGroup0.File("", "public/about.html")

	apiGroup1 := e.Group("/api/v1/projects")
	apiGroup1.GET("", api_projects.GetProjectsDataFromGithub)

	apiGroup2 := e.Group("/commands")
	apiGroup2.POST("", commands.CommandHandler)

	apiGroup3 := e.Group("")
	apiGroup3.File("", "public/home.html")

	apiGroup4 := e.Group("/project")
	apiGroup4.GET("", project.ProjectHandler)
}
