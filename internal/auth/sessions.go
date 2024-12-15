package auth

import (
	"encoding/gob"

	"github.com/bismastr/discord-bot/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
)

type SessionOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool
	Secure     bool
}

func NewSessionStore(opts SessionOptions) *sessions.CookieStore {
	var store = sessions.NewCookieStore([]byte(opts.CookiesKey))
	gob.Register(goth.User{})
	gob.Register(map[string]interface{}{})

	store.MaxAge(opts.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure
	store.Options.Domain = config.Envs.CookiesDomain

	return store
}
