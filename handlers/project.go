package handlers

import (
	"net/http"

	"github.com/harish876/portfolio/views/project"
	"github.com/labstack/echo/v4"
)

func ProjectHandler(c echo.Context) error {
	return Render(c, http.StatusOK, project.Project())
}
