package routes

import (
	"street-art/db"

	"github.com/gorilla/mux"
	auth "street-art/routes/middlewares"
)

func ApiRoutes(r *mux.Router, manager *db.Manager) {
	api := r.PathPrefix("/api").Subrouter()
	api.Use(auth.AuthRequired("user"))

}