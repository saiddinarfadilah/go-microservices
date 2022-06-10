package config

import "os"

const (
	secretgithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(secretgithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
