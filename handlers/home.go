package handlers

import (
	"net/http"

	"github.com/harish876/portfolio/views/home"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, home.Home())
}
