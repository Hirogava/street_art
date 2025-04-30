package handlers

import (
	"html/template"
	"log"
	"net/http"
	"street-art/db"
)

func Landing(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	categories, err := manager.GetAllCategories()
	if err != nil {
		log.Println("Ошибка получения категорий: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topProducts, err := manager.GetTopProducts()
	if err != nil {
		log.Println("Ошибка получения топ товаров: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"static/html/index.html",
		"static/html/templates/header.html",
		"static/html/templates/footer.html",
	))
	tmpl.Execute(w, map[string]interface{}{
		"Categories": categories,
		"Products":   topProducts,
	})
}