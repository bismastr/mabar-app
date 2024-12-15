package auth

import (
	"fmt"
	"net/http"

	"github.com/bismastr/discord-bot/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

type AuthService struct{}

const (
	SessionName = "_user_session"
)

func NewAuthService(store sessions.Store) *AuthService {
	gothic.Store = store

	goth.UseProviders(
		discord.New(
			config.Envs.DiscordClientID,
			config.Envs.DiscordClientSecret,
			buildCallbackURL("discord"),
			discord.ScopeIdentify,
			discord.ScopeEmail),
	)
	return &AuthService{}
}

func (a *AuthService) StoreUserSession(w http.ResponseWriter, r *http.Request, user goth.User) error {
	session, _ := gothic.Store.Get(r, SessionName)
	session.Values["user"] = user

	err := session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) GetUserSession(w http.ResponseWriter, r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return goth.User{}, fmt.Errorf("please login")
	}

	u := session.Values["user"]
	if u == nil {
		return goth.User{}, fmt.Errorf("user is not authenticated! %v", u)
	}

	return u.(goth.User), nil
}

func buildCallbackURL(provider string) string {
	if config.Envs.Env == "PRODUCTION" {
		return fmt.Sprintf("%s/api/v1/auth/%s/callback", config.Envs.PublicHost, provider)
	} else {
		return fmt.Sprintf("%s%s/api/v1/auth/%s/callback", config.Envs.PublicHost, config.Envs.Port, provider)
	}

}
