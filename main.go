package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/phihdn/projectmotor/auth"
	"github.com/phihdn/projectmotor/database"
	"github.com/phihdn/projectmotor/handler"
)

var store = sessions.NewCookieStore([]byte("0c276f3b2aa511358f443b7bd2188bbf68eba38ac2531f20772293f1d4a3ced0"))

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	err = database.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(handler.HandlerOptions{
		DB:    db,
		Store: store,
	})
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/login", h.Login)
	r.Get("/oauth/github/login", h.OAuthGitHubLogin)
	r.Get("/oauth/github/callback", h.OAuthGitHubCallback)

	// Group protected routes into one function to run middleware
	r.Group(protectedRouter(h))
	http.ListenAndServe("localhost:3000", r)
}

// Redirect to public auth route
//
// Use this when session user does not exist
func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:3000/login", http.StatusSeeOther)
}

// Protected context
//
// Middleware checks if user exists within current session
func ProtectedCtx(h *handler.Handler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			session, err := h.GetSessionStore(r)
			// redirect in case of error
			if err != nil {
				redirectToLogin(w, r)
				return
			}
			userID := session.Values["userID"]
			// redirect in case of missing user
			if userID == nil {
				redirectToLogin(w, r)
				return
			}
			// check if userID is of type int32
			ctx := r.Context()
			if userID, ok := userID.(int32); ok {
				user, _, err := h.UserService.GetUserByID(userID)
				if err != nil {
					redirectToLogin(w, r)
					return
				}
				ctx = context.WithValue(ctx, auth.UserKey{}, user)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// Router with user ensured
//
// Add routes here where user has to be logged in
func protectedRouter(h *handler.Handler) func(chi.Router) {
	return func(r chi.Router) {
		r.Use(ProtectedCtx(h))
		r.Get("/projects", h.GetProjects)
		r.Get("/", h.Dashboard)
	}
}
