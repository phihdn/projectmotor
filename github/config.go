package github

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var Config = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	Scopes:       []string{"read:user", "user:email"},
	Endpoint:     github.Endpoint,
}
