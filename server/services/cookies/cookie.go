package cookies

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type Manager struct {
	Session *sessions.Session
}

var store *sessions.CookieStore

func Init(key string) {
	store = sessions.NewCookieStore([]byte(key))
	store.Options.HttpOnly = true
	store.Options.Secure = false
	store.Options.SameSite = http.SameSiteStrictMode
}

func NewCookieManager(r *http.Request) *Manager {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Printf("Ошибка при получении сессии: %v", session)
	}
	return &Manager{
		Session: session,
	}
}