package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iamtbay/todo/controllers"
	"github.com/iamtbay/todo/helpers"
)

func Todos() http.Handler {

	r := chi.NewRouter()
	//
	r.Use(helpers.CheckJWTMiddleware)
	r.Use(helpers.LoggingMiddleware)
	//ADD MIDDLEWARE TO AUTH
	r.Get("/", controllers.GetTodos)
	r.Post("/", controllers.CreateNewTodo)
	r.Get("/{id}", controllers.GetSingleTodo)
	r.Patch("/{id}", controllers.UpdateATodo)
	r.Patch("/mark/{id}", controllers.MarkAsCompleted)
	r.Delete("/{id}", controllers.DeleteATodo)

	return r

}
