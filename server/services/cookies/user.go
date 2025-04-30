package cookies

import (
	"net/http"
	"street-art/models"
)

func SaveUserRoleCookie(r *http.Request, w http.ResponseWriter, role string, store *Manager) error {
	store.Session.Values["role"] = role

	err := store.Session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func SaveUserCookie(r *http.Request, w http.ResponseWriter, user models.User) error {
	store := NewCookieManager(r)

	store.Session.Values["id"] = user.Id
	store.Session.Values["name"] = user.Name
	store.Session.Values["email"] = user.Email
	store.Session.Values["phone"] = user.Phone
	store.Session.Values["address"] = user.Address
	store.Session.Values["balance"] = user.Balance

	if store.Session.Values["role"] != nil {
		return nil
	}
	return SaveUserRoleCookie(r, w, "user", store)
}

func GetUserCookie(r *http.Request, w http.ResponseWriter) models.User {
	var user models.User

	store := NewCookieManager(r)
	user.Id = store.Session.Values["id"].(int)
	user.Name = store.Session.Values["name"].(string)
	user.Email = store.Session.Values["email"].(string)
	user.Phone = store.Session.Values["phone"].(string)
	user.Address = store.Session.Values["address"].(string)
	user.Balance = store.Session.Values["balance"].(float64)

	return user
}

func EditUserCookie(r *http.Request, w http.ResponseWriter, user models.User) error {
	store := NewCookieManager(r)

	store.Session.Values["name"] = user.Name
	store.Session.Values["phone"] = user.Phone
	store.Session.Values["address"] = user.Address

	if err := store.Session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func RemoveUserCookie(r *http.Request, w http.ResponseWriter) error {
	store := NewCookieManager(r)
	store.Session.Options.MaxAge = -1
	if err := store.Session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func GetUserIdCookie(r *http.Request) int {
	store := NewCookieManager(r)
	return store.Session.Values["id"].(int)
}