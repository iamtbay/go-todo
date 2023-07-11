package helpers

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "CURRENT_USER")
		sessionUser := session.Values["user"]
		if sessionUser == nil {
			WriteJSON(w, http.StatusUnauthorized, ErrorJson{Error: "Unauthorized!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
//JWT CHECK

func CheckJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := GetCookie(r, "accessToken")
		if err != nil {
			WriteJSON(w, http.StatusBadGateway, ErrorJson{Error: "invalid token"})
			return
		}
		//validate jwt
		_, err = ValidateJWT(cookie.Value)
		if err != nil {
			log.Println("jwt err")
			return
		}
		next.ServeHTTP(w, r)
	})
}
