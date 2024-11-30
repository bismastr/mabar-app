package auth

import (
	"github.com/gorilla/sessions"
)

type SessionOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool
	Secure     bool
}

func NewSessionStore(opts SessionOptions) *sessions.CookieStore {
	var store = sessions.NewCookieStore([]byte(opts.CookiesKey))

	store.MaxAge(opts.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure
	store.Options.Domain = ".bism.app"

	return store
}
