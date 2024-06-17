package main

import (
	"fmt"
	"os"

	"github.com/harish876/portfolio/handlers"
	"github.com/harish876/portfolio/routes"
	"github.com/harish876/portfolio/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	if _, err := utils.GetEnv(); err != nil {
		fmt.Println(err)
	}

	app := echo.New()
	app.Static("/", "assets")

	app.HTTPErrorHandler = handlers.NotFoundErrorHandler
	routes.RegisterRoutes(app)
	// routes.RegisterHomeRoutes(app)
	// routes.RegisterAboutRoutes(app)
	// routes.RegisterCommandRoutes(app)
	// routes.RegisterProjectRoutes(app)
	// routes.RegisterProjectApiRoutes(app)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
