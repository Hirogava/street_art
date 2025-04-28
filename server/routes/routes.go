package routes

import (
	"net/http"
	"street-art/db"
	"street-art/handlers/pages"
	auth "street-art/routes/middlewares"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, manager *db.Manager) {
	Routes(r, manager)
	ApiRoutes(r, manager)
}

func Routes(r *mux.Router, manager *db.Manager) {
	user := r.PathPrefix("/profile").Subrouter()
	user.Use(auth.AuthRequired("user"))

	/*
		Роуты для авторизации
	*/
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(w, r)
	}).Methods(http.MethodGet)

	/*
		Роуты для страниц продуктов
	*/
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Landing(w, r, manager)
	}).Methods(http.MethodGet)

	r.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Product(w, r, manager)
	}).Methods(http.MethodGet)

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.Products(w, r, manager)
	}).Methods(http.MethodGet)

	r.HandleFunc("/products/{category_id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.ProductsByCategory(w, r)
	}).Methods(http.MethodGet)

	/*
		Роуты для страницы корзины
	*/
	user.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		handlers.Cart(w, r, manager)
	}).Methods(http.MethodGet)

	/*
		Роуты для профиля
	*/
	user.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		handlers.Profile(w, r, manager)
	}).Methods(http.MethodGet)

	user.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		handlers.EditProfile(w, r, manager)
	}).Methods(http.MethodGet)

	user.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.Orders(w, r, manager)
	}).Methods(http.MethodGet)

	user.HandleFunc("/order/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.OrderDetails(w, r, manager)
	}).Methods(http.MethodGet)
}