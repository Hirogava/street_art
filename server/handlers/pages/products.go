package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"street-art/db"
	"street-art/models"

	"github.com/gorilla/mux"
)

func Products(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	products, err := manager.GetAllProducts()
	if err != nil {
		log.Println("Ошибка получения всех продуктов: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/html/products.html"))
	tmpl.Execute(w, products)
}

func Product(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка конвертации id: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := manager.GetProductById(id)
	if err != nil {
		if err.Error() == "product not found" {
			http.NotFound(w, r)
		} else {
			log.Println("Ошибка получения товара: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	tmpl := template.Must(template.ParseFiles("static/html/product.html"))
	tmpl.Execute(w, product)
}

func ProductsByCategory(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		log.Println("Ошибка конвертации id: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := manager.GetAllProductsByCategoryId(categoryId)
	if err != nil {
		if err.Error() == "product not found" {
			http.NotFound(w, r)
		} else {
			log.Println("Ошибка получения товара: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	data := struct {
		Products []models.Product
		Category string
	}{
		Products: products,
		Category: products[0].Category,
	}

	tmpl := template.Must(template.ParseFiles("static/html/productsbycat.html"))
	tmpl.Execute(w, data)
}