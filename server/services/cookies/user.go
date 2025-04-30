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

func GetUserCookie(r *http.Request, w http.ResponseWriter) (*models.User, error) {
    user := models.User{}

    store := NewCookieManager(r)
    
    if store.Session == nil {
        return nil, nil
    }

    if id, ok := store.Session.Values["id"].(int); ok {
        user.Id = id
    } else {
        return nil, nil
    }

    if name, ok := store.Session.Values["name"].(string); ok {
        user.Name = name
    } else {
        user.Name = ""
    }

    if email, ok := store.Session.Values["email"].(string); ok {
        user.Email = email
    } else {
        user.Email = ""
    }

    if phone, ok := store.Session.Values["phone"].(string); ok {
        user.Phone = phone
    } else {
        user.Phone = ""
    }

    if address, ok := store.Session.Values["address"].(string); ok {
        user.Address = address
    } else {
        user.Address = ""
    }

    if balance, ok := store.Session.Values["balance"].(float64); ok {
        user.Balance = balance
    } else {
        user.Balance = 0.0
    }

    return &user, nil
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