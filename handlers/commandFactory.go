package handlers

import (
	"net/http"
	"strings"

	"github.com/harish876/portfolio/models"
	"github.com/harish876/portfolio/views/about"
	"github.com/harish876/portfolio/views/help"
	"github.com/harish876/portfolio/views/project"
	"github.com/labstack/echo/v4"
)

var (
	ABOUT    = "about"
	AUTHOR   = "author"
	PROJECT  = "project"
	PROJECTS = "projects"
	BLOG     = "blog"
	BLOGS    = "blogs"
	HOME     = "home"

	HELP = "help"
)

// TODO: move this to a different logical group
var (
	HELP_OPTIONS = []models.Commands{
		{Command: "about", Description: "check all available commands"},
		{Command: "readme", Description: "check out the guy who made this (using htmx btw)"},
		{Command: "man", Description: "check all keys only navigation options"},
		{Command: "project", Description: "check out all my projects"},
		{Command: "blog", Description: "check out all my blogs"},
		{Command: "cd ~/", Description: "go back to the home page"},
		{Command: "profile", Description: "go to my socials"},
	}
)

func CommandHandler(c echo.Context) error {
	command := c.FormValue("command")
	command = strings.ToLower(command)
	switch command {
	case ABOUT:
		return Render(c, http.StatusOK, about.About())
	case PROJECT, PROJECTS:
		return Render(c, http.StatusOK, project.Project())
	case HELP:
		return Render(c, http.StatusOK, help.Help(HELP_OPTIONS))
	default:
		//todo make a new one
		return Render(c, http.StatusOK, about.About())
	}
}
