package routes

import (
	"net/http"
	"street-art/db"
	"street-art/handlers/api"

	auth "street-art/routes/middlewares"

	"github.com/gorilla/mux"
)

func ApiRoutes(r *mux.Router, manager *db.Manager) {
	authRoutes := r.PathPrefix("/api").Subrouter()

	/*
		Роуты для входа
	*/
	authRoutes.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		api.Login(manager, w, r)
	}).Methods(http.MethodPost)

	authRoutes.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		api.Register(manager, w, r)
	}).Methods(http.MethodPost)

	apiRout := r.PathPrefix("/api").Subrouter()
	apiRout.Use(auth.AuthRequired("user"))

	apiRout.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		api.Logout(manager, w, r)
	}).Methods(http.MethodGet)

	/*
		Роуты для редактирования профиля
	*/
	apiRout.HandleFunc("/edit_profile", func(w http.ResponseWriter, r *http.Request) {
		api.EditProfile(manager, w, r)
	}).Methods(http.MethodPost)
}