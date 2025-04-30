package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"street-art/db"
	"street-art/file-manager"
	"street-art/models"

	"github.com/gorilla/mux"
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

func DeleteProduct(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка парсинга ID: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	imageUrl, err := manager.DeleteProduct(id)
	if err != nil {
		log.Println("Ошибка удаления продукта из бд: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	err = filemanager.DeletePhoto(imageUrl)
	if err != nil {
		log.Println("Ошибка удаления картинки продукта: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	var product models.Product
	var err error
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("Ошибка парсинга формы: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	vars := mux.Vars(r)
	product.Id, err = strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка парсинга ID: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.Name = r.FormValue("name")
	product.Description = r.FormValue("description")
	product.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Ошибка парсинга цены: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.Stock, err = strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		log.Println("Ошибка парсинга количества: " + err.Error())
  		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.CategoryId, err = strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		log.Println("Ошибка парсинга ID категории: " + err.Error())
  		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.BrandId, err = strconv.Atoi(r.FormValue("brand_id"))
	if err != nil {
		log.Println("Ошибка парсинга ID бренда: " + err.Error())
  		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	product.ImageUrl = r.FormValue("image_url")
	if files := r.MultipartForm.File["image"]; len(files) > 0 {
		image, imageHeader, err := r.FormFile("image")
		if err != nil {
			log.Println("Ошибка чтения изображения: " + err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
			return
		}
		defer image.Close()
	
		newImageUrl, err := filemanager.UpdateProductPhoto(image, imageHeader, product.CategoryId, product.ImageUrl)
		if err != nil {
			log.Println("Ошибка обновления изображения: " + err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
			return
		}
		product.ImageUrl = newImageUrl
	}

	err = manager.UpdateProduct(&product)
	if err != nil {
		log.Println("Ошибка обновления продукта: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})
}