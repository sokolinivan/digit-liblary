package app

import (
	"context"
	models "digit-liblary/models"
	u "digit-liblary/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowAnon := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range allowAnon {
			if value == requestPath {
				next.ServeHTTP(w, r)
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authrization")

		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("content-type", "application/json")
			u.Respond(w, response)
			
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := u.Message(false, "Invalid token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("content-type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1];
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func (token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_key")), nil
		})

		if err != nil {
			response = u.Message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("content-type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("content-type", "application/json")
			u.Respond(w, response)
			return
		}

		fmt.Sprintf("User %s", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}