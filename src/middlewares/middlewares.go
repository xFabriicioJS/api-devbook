package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate verifies if the requesting user is already authenticated or not
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
