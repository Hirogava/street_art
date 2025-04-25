package cookies

import (
	"net/http"
	"street-art/models"
)

func SaveUserCookie(r *http.Request, w http.ResponseWriter, user models.User) error {
	store := NewCookieManager(r)

	store.Session.Values["id"] = user.Id
	store.Session.Values["name"] = user.Name
	store.Session.Values["email"] = user.Email
	store.Session.Values["phone"] = user.Phone
	store.Session.Values["address"] = user.Address

	return store.Session.Save(r, w)
}

func GetUserCookie(r *http.Request, w http.ResponseWriter) models.User {
	var user models.User

	store := NewCookieManager(r)
	user.Id = store.Session.Values["id"].(int)
	user.Name = store.Session.Values["name"].(string)
	user.Email = store.Session.Values["email"].(string)
	user.Phone = store.Session.Values["phone"].(string)
	user.Address = store.Session.Values["address"].(string)

	return user
}

func EditUserCookie(r *http.Request, w http.ResponseWriter, user models.User) error {
	store := NewCookieManager(r)

	store.Session.Values["name"] = user.Name
	store.Session.Values["phone"] = user.Phone
	store.Session.Values["address"] = user.Address

	return store.Session.Save(r, w)
}

func RemoveUserCookie(r *http.Request, w http.ResponseWriter) error {
	store := NewCookieManager(r)
	store.Session.Options.MaxAge = -1
	if err := store.Session.Save(r, w); err != nil {
		return err
	}
	return nil
}