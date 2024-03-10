package handler

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/phihdn/projectmotor/template"
)

func (h Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	user := h.GetUserFromContext(r.Context())
	projects, err := h.ProjectService.GetAll(user.ID)
	if err != nil {
		fail(w, err, http.StatusInternalServerError)
		return
	}
	err = template.Projects.ExecuteWriter(pongo2.Context{
		"Projects": projects,
	}, w)
	if err != nil {
		fail(w, err, http.StatusInternalServerError)
		return
	}
}
