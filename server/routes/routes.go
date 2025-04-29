package routes

import (
	"net/http"
	"street-art/db"
	"street-art/handlers/pages"
	panel "street-art/handlers/admin"
	auth "street-art/routes/middlewares"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, manager *db.Manager) {
	Routes(r, manager)
	ApiRoutes(r, manager)
	AdminRoutes(r, manager)
	AdminApiRoutes(r, manager)
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
		handlers.ProductsByCategory(w, r, manager)
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

func AdminRoutes(r *mux.Router, manager *db.Manager) {
	admin := r.PathPrefix("/admin").Subrouter()
	adminWithoutAuth := r.PathPrefix("/admin").Subrouter()
	admin.Use(auth.AuthRequired("admin"))

	adminWithoutAuth.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		panel.Login(w, r, manager)
	}).Methods(http.MethodGet)

	admin.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		panel.Dashboard(w, r, manager)
	})

	admin.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		panel.Users(w, r, manager)
	})

	admin.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		panel.User(w, r, manager)
	})

	admin.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		panel.Products(w, r, manager)
	})

	admin.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		panel.Product(w, r, manager)
	})

	admin.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		panel.Orders(w, r, manager)
	})

	admin.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		panel.Order(w, r, manager)
	})

	admin.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		panel.Categories(w, r, manager)
	})

	admin.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		panel.Category(w, r, manager)
	})
}