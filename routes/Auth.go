package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iamtbay/todo/controllers"
)

func Auth() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", controllers.Login)
	r.Post("/register", controllers.Register)
	//AUTH MIDDLEWARE

	r.Post("/logout", controllers.Logout)
	return r
}
