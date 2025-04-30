package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"street-art/db"
	"street-art/file-manager"
	"street-art/models"
)

func AddProduct(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}
	var product models.Product
	var err error

	product.Name = r.FormValue("name")
    product.Description = r.FormValue("description")
	product.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.Stock, err = strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.CategoryId, err = strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.BrandId, err = strconv.Atoi(r.FormValue("brand_id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	image, fileHeader, err := r.FormFile("image")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}
	defer image.Close()

	product.ImageUrl, err = filemanager.SaveProductPhoto(image, fileHeader, product.CategoryId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	err = manager.AddProduct(&product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}