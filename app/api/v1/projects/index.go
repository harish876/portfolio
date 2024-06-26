package api_projects

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/harish876/portfolio/models"
	"github.com/harish876/portfolio/utils"
	"github.com/labstack/echo/v4"
	"github.com/machinebox/graphql"
)

func GetProjectsDataFromGithubHandler() (*models.PinnedItems, error) {
	client := graphql.NewClient("https://api.github.com/graphql")
	envVars, err := utils.GetEnv()
	var githubToken string
	if err != nil || envVars == nil {
		githubToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	} else {
		githubToken = envVars.GithubAccessToken
	}

	slog.Debug("GetProjectsDataFromGithubHandler", "github token", githubToken)

	req := graphql.NewRequest(`
        query getData($name: String!, $pinned: Int!, $languages: Int!) {
            user(login: $name) {
                pinnedItems(first: $pinned) {
                    totalCount
                    edges {
                        node {
                            ... on Repository {
                                name
                                id
                                url
                                description
                                homepageUrl
                                languages(first: $languages) {
                                    edges {
                                        node {
                                            id
                                            name
                                        }
                                    }
                                }
                                updatedAt
                            }
                        }
                    }
                }
            }
        }
    `)

	req.Var("name", "harish876")
	req.Var("pinned", 6)
	req.Var("languages", 2)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubToken))

	ctx := context.Background()

	var respData models.ResponseData
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return &respData.User.PinnedItems, nil
}

// @get
func GetProjectsDataFromGithub(c echo.Context) error {
	respData, err := GetProjectsDataFromGithubHandler()
	if err != nil {
		return c.String(http.StatusOK, "Failed to fetch data. Token Expired")
	}
	return c.JSON(http.StatusOK, respData)
}
