package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/harish876/portfolio/handlers"
	"github.com/harish876/portfolio/routes"
	"github.com/harish876/portfolio/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	if _, err := utils.GetEnv(); err != nil {
		slog.Error("Unable to load .env file", "error", err)
	}

	app := echo.New()
	slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	app.Static("/", "assets")

	app.HTTPErrorHandler = handlers.NotFoundErrorHandler
	routes.RegisterRoutes(app)
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", (os.Getenv("PORT")))))
}
