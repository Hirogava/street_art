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
	apiProfile := apiRout.PathPrefix("/profile").Subrouter()

	apiProfile.HandleFunc("/edit_profile", func(w http.ResponseWriter, r *http.Request) {
		api.EditProfile(manager, w, r)
	}).Methods(http.MethodPost)

	/*
		Роуты для товаров
	*/
	apiProduct := apiRout.PathPrefix("/product").Subrouter()

	apiProduct.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		api.AddToCart(w, r, manager)
	}).Methods(http.MethodPost)

	apiProduct.HandleFunc("/cart/{id}/clear", func(w http.ResponseWriter, r *http.Request) {
		api.ClearCart(w, r, manager)
	}).Methods(http.MethodPost)

	apiProduct.HandleFunc("/cart/{cart_id}/delete/{product_id}", func(w http.ResponseWriter, r *http.Request) {
		api.DeleteProductFromCart(w, r, manager)
	}).Methods(http.MethodPost)

	/*
		Роуты для денежных операций
	*/
	apiBalance := apiRout.PathPrefix("/balance").Subrouter()
	
	apiBalance.HandleFunc("/deposit", func(w http.ResponseWriter, r *http.Request) {
		api.Deposit(w, r, manager)
	}).Methods(http.MethodPost)

	apiBalance.HandleFunc("/purchase", func(w http.ResponseWriter, r *http.Request) {
		api.Purchase(w, r, manager)
	}).Methods(http.MethodPost)
}

func AdminApiRoutes(r *mux.Router, manager *db.Manager) {
	adminRout := r.PathPrefix("/api/admin").Subrouter()
	adminRoutWithoutAuth := r.PathPrefix("/api/admin").Subrouter()
	adminRout.Use(auth.AdminRequired("admin"))

	adminRoutWithoutAuth.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		api.AdminLogin(w, r, manager)
	}).Methods(http.MethodPost)

	adminRout.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		api.AdminLogout(w, r, manager)
	})

	adminRout.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		api.GetCategoriesForPanel(w, r, manager)
	}).Methods(http.MethodGet)

	adminRout.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		api.AddCategory(w, r, manager)
	}).Methods(http.MethodPost)

	adminRout.HandleFunc("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.DeleteCategory(w, r, manager)
	}).Methods(http.MethodDelete)

	adminRout.HandleFunc("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.UpdateCategory(w, r, manager)
	}).Methods(http.MethodPut)

	adminRout.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		api.AddProduct(w, r, manager)
	}).Methods(http.MethodPost)

	adminRout.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.DeleteProduct(w, r, manager)
	}).Methods(http.MethodDelete)

	adminRout.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.UpdateProduct(w, r, manager)
	}).Methods(http.MethodPut)
}