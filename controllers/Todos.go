package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/iamtbay/todo/database"
	"github.com/iamtbay/todo/helpers"
)

// DECLARE IT TO USE ALL FUNCTIONS AS TODOS.
var todos *database.Todo
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// CREATE NEW TODO
func CreateNewTodo(w http.ResponseWriter, r *http.Request) {
	//
	var newTodo *database.Todo
	cookie, err := helpers.GetCookie(r, "accessToken")
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadGateway, helpers.ErrorJson{Error: "invalid token"})
		return
	}
	//validate jwt
	_, err = helpers.ValidateJWT(cookie.Value)
	if err != nil {
		log.Println("jwt err")
		return
	}
	//DECODE THE USER'S VALUE
	err = json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: "error while decoding"})
		fmt.Println(err)
		return
	}
	//GET ACTIVE USER'S ID TO CREATE A NEW TODO
	session, _ := store.Get(r, "CURRENT_USER")
	currentUser := helpers.GetUser(session)
	newTodo.UserId = currentUser.Id

	//CREATE TODO
	err = todos.CreateNewTodo(newTodo)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: "error while creating todo"})
		fmt.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK,
		struct {
			Msg string `json:"msg"`
		}{Msg: "Todo Created"})
}

// ---
// GET USER'S TODOS
func GetTodos(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "CURRENT_USER")
	currentUser := helpers.GetUser(session)

	todos, err := todos.GetTodos(currentUser.Id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: "error while getting todos"})
		fmt.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, todos)

}

//GET SINGLE TODO

func GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	//GET URL ID
	id := chi.URLParam(r, "id")
	//GET SESSIONS
	sessions, _ := store.Get(r, "CURRENT_USER")
	currentUser := helpers.GetUser(sessions)
	//DB OPERATIONS
	getSingleRow, err := todos.GetSingleTodo(id, currentUser.Id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
	}

	helpers.WriteJSON(w, http.StatusOK, getSingleRow)
}

// UPDATE A TODO
func UpdateATodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//get active user id
	sessions, _ := store.Get(r, "CURRENT_USER")
	currentUser := helpers.GetUser(sessions)
	//update operations
	//get user entry infos
	var infos *database.Todo
	err := json.NewDecoder(r.Body).Decode(&infos)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}
	data, err := todos.UpdateATodo(infos, id, currentUser.Id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}
	fmt.Println(data)

	helpers.WriteJSON(w, http.StatusOK, struct {
		Msg string `json:"message"`
	}{Msg: "updated!"})

}

// DELETE A TODO
func DeleteATodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sessions, _ := store.Get(r, "CURRENT_USER")
	currentUser := helpers.GetUser(sessions)

	err := todos.DeleteATodo(id, currentUser.Id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, struct {
		Msg string `json:"msg"`
	}{Msg: "succesfully deleted"})

}

// MARK AS COMPLETED

func MarkAsCompleted(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sessions, _ := store.Get(r, "CURRENT_USER")
	currentUser := helpers.GetUser(sessions)
	err := todos.MarkAsCompleted(id, currentUser.Id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, struct {
		Msg string `json:"msg"`
	}{Msg: "succesfully marked"})
}
