package handler

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/phihdn/projectmotor/template"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var config = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	Scopes:       []string{"read:user", "user:email"},
	Endpoint:     github.Endpoint,
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	template.Login.ExecuteWriter(pongo2.Context{}, w)
}
