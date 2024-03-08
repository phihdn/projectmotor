package handler

import (
	"crypto/rand"
	"encoding/hex"
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

func (h *Handler) OAuthGitHubLogin(w http.ResponseWriter, r *http.Request) {
	state, err := generateCSRFToken(16)
	if err != nil {
		fail(w, err, http.StatusInternalServerError)
		return
	}
	session, err := h.GetSessionStore(r)
	if err != nil {
		fail(w, err, http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		fail(w, err, http.StatusInternalServerError)
		return
	}
	url := config.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func generateCSRFToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
