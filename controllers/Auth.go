package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iamtbay/todo/database"
	"github.com/iamtbay/todo/helpers"
)

type Auth struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type MsgJSON struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

var UserStr *database.User

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {

	//CHECK USER EXIST OR NOT IN SESSIONS
	err := helpers.CheckExistUser(r)
	//if user exist return an error response
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}

	//GET BODY VALUES
	var user *database.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user, err = UserStr.Login(user)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}
	//SET USER'S COOKIE,JWT ETC. TO MAKE A OPERATIONS.

	//GENERATE TOKEN
	token, err := helpers.CreateJWT(user)
	if err != nil {
		panic(err)
	}
	// set cookie for storing token
	helpers.CreateCookie(w, "accessToken", token)
	//create session for user
	err = helpers.SetUserSession(w, r, user)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, Auth{Name: user.Name, Email: user.Email})

}

// REGISTER
func Register(w http.ResponseWriter, r *http.Request) {
	var user *database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("register decoder error")
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: err.Error()})
		return
	}

	err = user.Register(user)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.ErrorJson{Error: "Invalid credentials!"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, MsgJSON{Message: "User created!", Success: true})

}


//LOGOUT
func Logout(w http.ResponseWriter, r *http.Request) {
	err := helpers.UserAuth(r)
	if err != nil {
		helpers.WriteJSON(w, http.StatusUnauthorized, helpers.ErrorJson{Error: err.Error()})
		return
	}

	err = helpers.DeleteUserSession(w, r)
	if err != nil {
		helpers.WriteJSON(w, http.StatusUnauthorized, helpers.ErrorJson{Error: err.Error()})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, MsgJSON{Message: "Logout succesfully", Success: true})
}

//FORGET PASSWORD
