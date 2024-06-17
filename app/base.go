package app

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/harish876/portfolio/components/notfound"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func NotFoundErrorHandler(err error, c echo.Context) {
	if h, ok := err.(*echo.HTTPError); ok {
		if h.Code == http.StatusNotFound {
			Render(c, http.StatusNotFound, notfound.NotFound())
		}
	}
}
