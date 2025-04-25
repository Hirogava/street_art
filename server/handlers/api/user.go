package api

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = cookies.SaveUserCookie(r, w, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = cookies.SaveUserCookie(r, w, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешная регистрация"})
}

func Logout(manager *db.Manager, w http.ResponseWriter, r *http.Request) {
	err := cookies.RemoveUserCookie(r, w)
	if err != nil {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = cookies.EditUserCookie(r, w, user)
	if err != nil {	
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Успешное редактирование профиля"})
}