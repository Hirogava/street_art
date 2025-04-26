package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"street-art/db"
	"street-art/models"
	"street-art/services/cookies"

	"github.com/gorilla/mux"
)

func AddToCart(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	var product models.ProductToCart

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Ошибка формата json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.UserId = cookies.GetUserIdCookie(r)
	err := manager.AddToCart(&product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func ClearCart(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка получения id товара", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = manager.ClearCart(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func DeleteProductFromCart(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	cartId, err := strconv.Atoi(vars["cart_id"])
	if err != nil {
		log.Println("Ошибка получения id товара", err)
		w.WriteHeader(http.StatusBadRequest)
		return
  	}	

	productId, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		log.Println("Ошибка получения id товара", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = manager.DeleteProductFromCart(cartId, productId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}