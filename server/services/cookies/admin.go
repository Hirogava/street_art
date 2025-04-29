package cookies

import (
	"net/http"
	"street-art/models"
)

func SaveAdminCookie(r *http.Request, w http.ResponseWriter, user models.Admin) error {
	store := NewCookieManager(r)

	store.Session.Values["id"] = user.Id
	store.Session.Values["email"] = user.Email

	if store.Session.Values["role"] != nil {
		return nil
	}
	return SaveUserRoleCookie(r, w, "admin", store)
}