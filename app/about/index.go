package about

import (
	"net/http"

	"github.com/harish876/portfolio/app"
	"github.com/harish876/portfolio/views/about"
	"github.com/labstack/echo/v4"
)

// @get
func AboutHandler(c echo.Context) error {
	return app.Render(c, http.StatusOK, about.About())
}
