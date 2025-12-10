package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("auth: %d %s", e.Code, e.Message)
}

func ValidateToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", &Error{401, "Missing Authorization header"}
	}
	token = strings.TrimSpace(token)
	if !strings.HasPrefix(token, "Bearer ") {
		return "", &Error{401, "Invalid token format"}
	}
	return token[7:], nil
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := ValidateToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenString, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("secretkey"), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, ok := tokenString.Claims.(jwt.MapClaims)
		if !ok || !claims.Valid() {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims.ValidTime() == false {
			http.Error(w, "Token expired or invalid", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func (c *Claims) ValidTime() bool {
	// get expiration time
	// compare to current time
	// return true if valid, false otherwise
	return time.Unix(c.ExpiresAt, 0).After(time.Now().UTC())
}