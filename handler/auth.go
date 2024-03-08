package handler

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/phihdn/projectmotor/template"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	template.Login.ExecuteWriter(pongo2.Context{}, w)
}
