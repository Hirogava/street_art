package routes

import (
	"net/http"
	"street-art/db"
	"street-art/handlers/api"

	auth "street-art/routes/middlewares"

	"github.com/gorilla/mux"
)

func ApiRoutes(r *mux.Router, manager *db.Manager) {
	apiRout := r.PathPrefix("/api").Subrouter()
	apiRout.Use(auth.AuthRequired("user"))

	apiRout.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		api.Login(manager, w, r)
	}).Methods(http.MethodPost)

	apiRout.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		api.Register(manager, w, r)
	}).Methods(http.MethodPost)

	apiRout.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		api.Logout(manager, w, r)
	}).Methods(http.MethodPost)

	apiRout.HandleFunc("/edit_profile", func(w http.ResponseWriter, r *http.Request) {
		api.EditProfile(manager, w, r)
	}).Methods(http.MethodPost)
}