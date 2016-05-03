package models

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	GithubId    string `json:"-"`
	githubToken string
}

var client *oauth2.Config

// Configure oauth client
func init() {
	clientId := "get from yaml file"
	clientSecret := "also get from yaml file"
	client = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		// TODO: Double check that these are the scopes that are needed
		Scopes: []string{
			"user:email",
			"public_repo",
		},
		Endpoint: github.Endpoint,
	}
}

func GetOauthUrl() string {
	return client.AuthCodeURL("test-state", oauth2.AccessTypeOnline)
}
