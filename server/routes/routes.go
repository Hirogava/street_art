package routes

import (
	"net/http"
	"street-art/db"
	"street-art/handlers"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, manager *db.Manager) {
	StaticRoutes(r, manager)
}

func StaticRoutes(r *mux.Router, manager *db.Manager) {


	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Landing(w, r)
	})

	r.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Product(w, r)
	})

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.Products(w, r)
	})
}