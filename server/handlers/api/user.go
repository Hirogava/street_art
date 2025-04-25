package api

import (
	"encoding/json"
	"log"
	"net/http"
	"street-art/db"
	"street-art/models"
	"street-art/services/cookies"
)

func Login(manager *db.Manager, w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := manager.Login(email, password)
	if err != nil {
		log.Println("Ошибка авторизации: " + err.Error())
		http.Error(w, `{"message": "Ошибка авторизации"}`, http.StatusInternalServerError)
		return
	}

	err = cookies.SaveUserCookie(r, w, user)
	if err != nil {
		log.Println("Ошибка сохранения cookie" + err.Error())
		http.Error(w, `{"message": "Ошибка сохранения cookie"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешная авторизация"})
}

func Register(manager *db.Manager, w http.ResponseWriter, r *http.Request) {
	var user models.User

	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	password := r.FormValue("password")
	user.Phone = r.FormValue("phone")
	user.Address = r.FormValue("address")

	err := manager.Register(&user, password)
	if err != nil {
		log.Println("Ошибка регистрации" + err.Error())
		http.Error(w, `{"message": "Ошибка регистрации"}`, http.StatusInternalServerError)
		return
	}

	err = cookies.SaveUserCookie(r, w, user)
	if err != nil {
		log.Println("Ошибка сохранения cookie" + err.Error())
		http.Error(w, `{"message": "Ошибка сохранения cookie"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешная регистрация"})
}

func Logout(manager *db.Manager, w http.ResponseWriter, r *http.Request) {
	err := cookies.RemoveUserCookie(r, w)
	if err != nil {
		log.Println("Ошибка удаления cookie" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditProfile(manager *db.Manager, w http.ResponseWriter, r *http.Request) {
	user := cookies.GetUserCookie(r, w)
	
	user.Name = r.FormValue("name")
	user.Phone = r.FormValue("phone")
	user.Address = r.FormValue("address")

	err := manager.UpdateUserInfo(&user)
	if err != nil {
		log.Println("Ошибка обновления данных пользователя" + err.Error())
		http.Error(w, `{"message": "Ошибка обновления данных пользователя"}`, http.StatusInternalServerError)
		return
	}

	err = cookies.EditUserCookie(r, w, user)
	if err != nil {	
		log.Println("Ошибка сохранения cookie" + err.Error())
		http.Error(w, `{"message": "Ошибка сохранения cookie"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешное редактирование профиля"})
}