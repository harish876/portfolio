package utils

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	GithubAccessToken string `envconfig:"GITHUB_ACCESS_TOKEN"`
}

var (
	envVars *Env
)

func GetEnv() (*Env, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}
	if envVars != nil {
		return envVars, nil
	}
	var s Env
	err := envconfig.Process("", &s)
	if err != nil {
		return nil, err
	}
	envVars = &s
	return envVars, nil

}
