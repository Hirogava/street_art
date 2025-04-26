package api

import (
	"encoding/json"
	"log"
	"net/http"
	"street-art/db"
	"street-art/models"
	"street-art/services/cookies"
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