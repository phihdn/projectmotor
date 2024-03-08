package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/phihdn/projectmotor/database"
)

type Handler struct {
	userService *database.UserService
	store       *sessions.CookieStore
}

type HandlerOptions struct {
	DB    *sqlx.DB
	Store *sessions.CookieStore
}

func NewHandler(options HandlerOptions) *Handler {
	userService := database.NewUserService(options.DB)
	return &Handler{
		userService: userService,
		store:       options.Store,
	}
}

func (h *Handler) GetSessionStore(r *http.Request) (*sessions.Session, error) {
	return h.store.Get(r, "_projectmotor_session")
}
