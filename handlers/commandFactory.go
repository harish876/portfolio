package handlers

import (
	"net/http"
	"strings"

	"github.com/harish876/portfolio/models"
	"github.com/harish876/portfolio/views/about"
	"github.com/harish876/portfolio/views/help"
	"github.com/harish876/portfolio/views/home"
	"github.com/harish876/portfolio/views/project"
	"github.com/labstack/echo/v4"
)

var (
	README   = "readme"
	PROJECT  = "project"
	PROJECTS = "projects"
	BLOG     = "blog"
	BLOGS    = "blogs"
	HOME     = "cd ~/"

	HELP = "help"
)

// TODO: move this to a different logical group
var (
	HELP_OPTIONS = []models.Commands{
		{Command: "readme", Description: "check out the guy who made this (using htmx btw)"},
		{Command: "man", Description: "check all keys only navigation options (coming soon)"},
		{Command: "project", Description: "check out all my projects"},
		{Command: "blog", Description: "check out all my blogs (coming soon)"},
		{Command: "cd ~/", Description: "go back to the home page"},
		{Command: "profile", Description: "go to my socials (coming soon)"},
	}
)

func CommandHandler(c echo.Context) error {
	command := c.FormValue("command")
	command = strings.ToLower(command)
	switch command {
	case HOME:
		return Render(c, http.StatusOK, home.Home())
	case README:
		return Render(c, http.StatusOK, about.About())
	case PROJECT, PROJECTS:
		data, _ := GetProjectsDataFromGithubHandler()
		return Render(c, http.StatusOK, project.Project(data))
	case HELP:
		return Render(c, http.StatusOK, help.Help(HELP_OPTIONS))
	default:
		return Render(c, http.StatusOK, home.Home())
	}
}
