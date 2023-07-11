package helpers

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/iamtbay/todo/database"
)

// SET USER SESSION
type ISessionUser struct {
	Name  string
	Email string
	Id    string
}

// SESSION SETTER
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// check exist user
func CheckExistUser(r *http.Request) error {
	//CHECK CURRENTLY USER
	//it's always returning  a session even if empty
	session, _ := store.Get(r, "CURRENT_USER")
	//check user in sessions
	sessionUser := session.Values["user"]
	// if it's empty
	if sessionUser != nil {
		return errors.New("please logout first")
	}
	return nil
}

// CHECK USER LOGGED IN OR NOT
func UserAuth(r *http.Request) error {
	session, _ := store.Get(r, "CURRENT_USER")
	sessionUser := session.Values["user"]
	if sessionUser == nil {
		fmt.Println(sessionUser)
		return errors.New("have to login")
	}
	return nil
}

// delete user session
func DeleteUserSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "CURRENT_USER")
	sessionExist := session.Values["user"]
	if sessionExist == nil {
		return errors.New("user couldn't find")
	}
	session.Values["user"] = nil
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

// SET USER SESSIONS
func SetUserSession(w http.ResponseWriter, r *http.Request, userInfos *database.User) error {
	//ITS REGISTER STRUCT TO USE IT ON SESSIONS
	gob.Register(ISessionUser{})
	gob.Register(map[string]ISessionUser{})
	//get session store
	session, _ := store.Get(r, "CURRENT_USER")
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,
		HttpOnly: true,
	}

	//create user session
	session.Values["user"] = ISessionUser{
		Name:  userInfos.Name,
		Email: userInfos.Email,
		Id:    userInfos.Id,
	}
	//save session
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//UNDERSTAND, HOW ITS GETTING VALUES?
func GetUser(s *sessions.Session) ISessionUser {
	val := s.Values["user"]
	var user = ISessionUser{}
	user, ok := val.(ISessionUser)
	if !ok {
		return ISessionUser{}
	}
	return user

}
