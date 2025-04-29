package api

import (
	"encoding/json"
	"log"
	"net/http"
	"street-art/db"
	"street-art/services/cookies"
)

func AdminLogin(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	admin, err := manager.AdminLogin(email, password)
	if err != nil {
		log.Println("Ошибка авторизации: " + err.Error())
		http.Error(w, `{"message": "Ошибка авторизации"}`, http.StatusInternalServerError)
		return
	}

	err = cookies.SaveAdminCookie(r, w, admin)
	if err != nil {
		log.Println("Ошибка сохранения cookie" + err.Error())
		http.Error(w, `{"message": "Ошибка сохранения cookie"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешная авторизация"})
}

func AdminLogout(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	err := cookies.RemoveUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка удаления cookie" + err.Error())
		http.Error(w, `{"message": "Ошибка выхода"}`, http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}