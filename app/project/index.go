package project

import (
	"net/http"

	"github.com/harish876/portfolio/app"
	api_projects "github.com/harish876/portfolio/app/api/v1/projects"
	"github.com/harish876/portfolio/views/project"
	"github.com/labstack/echo/v4"
)

// @get
func ProjectHandler(c echo.Context) error {
	data, _ := api_projects.GetProjectsDataFromGithubHandler()
	return app.Render(c, http.StatusOK, project.Project(data))
}
