package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"street-art/db"
	"street-art/services/cookies"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/register.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"Categories": categories, "User" : nil})
}

func Login(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/login.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"Categories": categories, "User" : nil})
}

func Profile(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	user, err := cookies.GetUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	balance, err := manager.GetUserBalance(user.Id)
	if err != nil {
		log.Println("Ошибка получения баланса пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Balance = balance

	tmpl := template.Must(template.ParseFiles(
		"static/html/profile.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"User": user, "Categories": categories})
}

func EditProfile(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	user, err := cookies.GetUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	balance, err := manager.GetUserBalance(user.Id)
	if err != nil {
		log.Println("Ошибка получения баланса пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Balance = balance

	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/edit_profile.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"User": user, "Categories": categories})
}

func Orders(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	userId := cookies.GetUserIdCookie(r)

	orders, err := manager.GetAllOrders(userId)
	if err != nil {
		if err.Error() == "no orders" {
			tmpl := template.Must(template.ParseFiles("static/html/orders.html",
				"static/html/templates/footer.html"))
			tmpl.Execute(w, nil)
			return
		}
		log.Println("Ошибка получения заказов пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := cookies.GetUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/orders.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"Orders": orders, "Categories": categories, "User": user})
}

func OrderDetails(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка при получении id заказа:", err)
		http.Error(w, "Ошибка получения id заказа: "+err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := manager.GetOrderDetails(orderId)
	if err != nil {
		log.Println("Ошибка получения деталей заказа:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := cookies.GetUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/order.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"Products": products, "Categories": categories, "User": user})
}

func Cart(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	userId := cookies.GetUserIdCookie(r)

	products, err := manager.GetAllCart(userId)
	if err != nil {
		if err.Error() == "empty" {
			tmpl := template.Must(template.ParseFiles("static/html/cart.html"))
			tmpl.Execute(w, nil)
			return
		}
		log.Println("Ошибка получения корзины пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := cookies.GetUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/cart.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{"Products": products, "Categories": categories, "User": user})
}