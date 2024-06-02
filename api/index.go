package handler

import (
	"net/http"

	"github.com/harish876/portfolio/handlers"
	"github.com/harish876/portfolio/routes"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// if _, err := utils.GetEnv(); err != nil {
	// 	fmt.Println("Unable to get env file")
	// 	//log.Fatal(err)
	// }

	app := echo.New()
	app.Static("/", "assets")

	app.HTTPErrorHandler = handlers.NotFoundErrorHandler
	routes.RegisterHomeRoutes(app)
	routes.RegisterAboutRoutes(app)
	routes.RegisterCommandRoutes(app)
	routes.RegisterProjectRoutes(app)
	routes.RegisterProjectApiRoutes(app)

	return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := Init()
	app.Logger.Fatal(app.Start(":42069"))
}
