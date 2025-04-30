package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"street-art/db"
	filemanager "street-art/file-manager"
	"street-art/models"

	"github.com/gorilla/mux"
)

func GetCategoriesForPanel(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	categories, err := manager.GetMiniCategories()
	if err != nil {
		if err.Error() == "no value" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "No categories"}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	brands, err := manager.GetMiniBrands()
	if err != nil {
		if err.Error() == "no value" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "No brands"}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	result := map[string]interface{}{
		"categories": categories,
		"brands":     brands,
	}

	jsonRes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func AddCategory(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}
	var category models.Category

	category.Name = r.FormValue("name")
	category.Description = r.FormValue("description")

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}
	defer image.Close()

	category.ImageUrl, err = filemanager.SaveCategoryPhoto(image, imageHeader)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	err = manager.AddCategory(&category)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)

	id, err  := strconv.Atoi(vars["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	filepath, err := manager.DeleteCategoryById(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
      	return
	}

	err = filemanager.DeletePhoto(filepath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Ошибка чтения параметра ID: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("Ошибка чтения формы: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	var category models.Category

	category.Id = id
	category.Name = r.FormValue("name")
	category.Description = r.FormValue("description")
	category.ImageUrl = r.FormValue("image_url")
	log.Println("Изображение:", category.ImageUrl)

	if files := r.MultipartForm.File["image"]; len(files) > 0 {
		image, imageHeader, err := r.FormFile("image")
		if err != nil {
			log.Println("Ошибка чтения изображения: " + err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
			return
		}
		defer image.Close()
	
		newImageUrl, err := filemanager.UpdateCategoryPhoto(image, imageHeader, category.ImageUrl)
		if err != nil {
			log.Println("Ошибка обновления изображения: " + err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
			return
		}
		category.ImageUrl = newImageUrl
		log.Println("Изображение обновлено:", newImageUrl)
	}

	err = manager.UpdateCategory(&category)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}