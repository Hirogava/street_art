package api

import (
	"encoding/json"
	"net/http"
	"street-art/db"
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
		w.Write([]byte(`{"message": "Error"}`))
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
		w.Write([]byte(`{"message": "Error"}`))
		return
	}

	result := map[string]interface{}{
		"categories": categories,
		"brands":     brands,
	}

	jsonRes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}