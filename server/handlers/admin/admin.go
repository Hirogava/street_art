package admin

import (
	"log"
	"net/http"
	"strconv"
	"street-art/db"
	"text/template"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	tmpl := template.Must(template.ParseFiles("static/html/admin/login.html"))
	tmpl.Execute(w, nil)
}

func Dashboard(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	tmpl := template.Must(template.ParseFiles("static/html/admin/dashboard.html"))
	tmpl.Execute(w, nil)
}

func Users(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	users, err := manager.GetAllUsers()
	if err != nil {
		log.Println("Ошибка получения списка пользователей:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/html/admin/users.html"))
	tmpl.Execute(w, users)
}

func Orders(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	orders, err := manager.GetAllOrdersForAdmin()
	if err != nil {
		if err.Error() == "no orders" {
			tmpl := template.Must(template.ParseFiles("static/html/admin/orders.html"))
			tmpl.Execute(w, nil)
			return
		}
		log.Println("Ошибка получения заказов пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/html/admin/orders.html"))
	tmpl.Execute(w, orders)
}

func Products(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	products, err := manager.GetAllProducts()
	if err != nil {
		if err.Error() == "no products" {
			tmpl := template.Must(template.ParseFiles("static/html/admin/products.html"))
			tmpl.Execute(w, nil)
			return
     	}
		log.Println("Ошибка получения списка продуктов:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("static/html/admin/products.html"))
	tmpl.Execute(w, products)
}

func Categories(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	categories, err := manager.GetAllCategories()
	if err != nil {
		if err.Error() == "no categories" {
			tmpl := template.Must(template.ParseFiles("static/html/admin/categories.html"))
			tmpl.Execute(w, nil)
			return
		}
		log.Println("Ошибка получения списка категорий:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/html/admin/categories.html"))
	tmpl.Execute(w, categories)
}

func Product(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка получения параметра id: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := manager.GetProductById(id)
	if err != nil {
		log.Println("Ошибка получения продукта:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/html/admin/product.html"))
	tmpl.Execute(w, product)
}

func Order(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	tmpl := template.Must(template.ParseFiles("static/html/admin/order.html"))
	tmpl.Execute(w, nil)
}

func Category(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	tmpl := template.Must(template.ParseFiles("static/html/admin/category.html"))
	tmpl.Execute(w, nil)
}