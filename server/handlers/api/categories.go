package api

import (
	"encoding/json"
	"net/http"
	"street-art/db"
	filemanager "street-art/file-manager"
	"street-art/models"
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