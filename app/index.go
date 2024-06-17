package app

import (
	"net/http"

	"github.com/harish876/portfolio/views/home"
	"github.com/labstack/echo/v4"
)

// @get
func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, home.Home())
}
