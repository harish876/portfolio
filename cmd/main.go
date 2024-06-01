package main

import (
	"github.com/harish876/portfolio/handlers"
	"github.com/harish876/portfolio/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")

	app.HTTPErrorHandler = handlers.NotFoundErrorHandler
	routes.RegisterHomeRoutes(app)
	routes.RegisterAboutRoutes(app)
	routes.RegisterCommandRoutes(app)
	routes.RegisterProjectRoutes(app)
	app.Logger.Fatal(app.Start(":42069"))
}
