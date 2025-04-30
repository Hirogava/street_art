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
            
            if store.Session == nil || store.Session.Values["role"] == nil {
                log.Println("Сессия не найдена")
                http.Redirect(w, r, "/register", http.StatusSeeOther)
                return
            }

            role, ok := store.Session.Values["role"].(string)
            if !ok {
                log.Println("Некорректный тип роли")
                http.Redirect(w, r, "/register", http.StatusSeeOther)
                return
            }

            if role != "user" {
                log.Printf("Неавторизованный пользователь: %s", role)
                http.Redirect(w, r, "/register", http.StatusSeeOther)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}

func AdminRequired(userType string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            store := cookies.NewCookieManager(r)
            
            if store.Session == nil || store.Session.Values["role"] == nil {
                log.Println("Сессия не найдена")
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            role, ok := store.Session.Values["role"].(string)
            if !ok {
                log.Println("Некорректный тип роли")
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            if role != "admin" {
                log.Printf("Неавторизованный пользователь: %s", role)
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}