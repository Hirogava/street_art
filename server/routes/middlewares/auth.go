package middleware

import (
	"street-art/services/cookies"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRequired(userType string) mux.MiddlewareFunc {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			store := cookies.NewCookieManager(r)
			if store.Session.Values["user_type"] != "user" {
				http.Redirect(w, r, "/register", http.StatusSeeOther)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func AdminRequired(userType string) mux.MiddlewareFunc {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			store := cookies.NewCookieManager(r)
			if store.Session.Values["user_type"] != "admin" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}