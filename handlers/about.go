package handlers

import (
	"net/http"

	"github.com/harish876/portfolio/views/about"
	"github.com/labstack/echo/v4"
)

func AboutHandler(c echo.Context) error {
	return Render(c, http.StatusOK, about.About())
}
