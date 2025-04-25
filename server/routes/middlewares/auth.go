package middleware

import (
	"log"
	"net/http"
	"street-art/services/cookies"

	"github.com/gorilla/mux"
)

func AuthRequired(userType string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			store := cookies.NewCookieManager(r)
			if store.Session.Values["role"].(string) != "user" {
				log.Println("Неавторизованный пользователь:" + store.Session.Values["role"].(string))
				http.Redirect(w, r, "/register", http.StatusSeeOther)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func AdminRequired(userType string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			store := cookies.NewCookieManager(r)
			if store.Session.Values["role"].(string) != "admin" {
				log.Println("Неавторизованный пользователь:" + store.Session.Values["role"].(string))
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
