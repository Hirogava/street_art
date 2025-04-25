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