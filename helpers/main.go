package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iamtbay/todo/database"
)

func CreateJWT(user *database.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"email":     user.Email,
		"id":        user.Id,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

//VALIDATE JWT

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(secret), nil
	})

}

// COOKIE CREATOR
func CreateCookie(w http.ResponseWriter, name, value string) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

}

// GET COOKIE
func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}

	return cookie, nil

}
