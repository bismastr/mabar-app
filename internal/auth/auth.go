package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/bismastr/discord-bot/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

type AuthService struct{}

const (
	SessionName = "user_session"
)

func NewAuthService(store sessions.Store) *AuthService {
	gothic.Store = store

	goth.UseProviders(
		discord.New(
			config.Envs.DiscordClientID,
			config.Envs.DiscordClientSecret,
			"http://localhost:8080/api/v1/auth/discord/callback",
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

func NewAuth() {
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET environment variable is not set")
	}

	var store = sessions.NewCookieStore([]byte(sessionSecret))

	gothic.Store = store

	discordKey := os.Getenv("DISCORD_KEY")
	discordSecret := os.Getenv("DISCORD_SECRET")
	if discordKey == "" || discordSecret == "" {
		log.Fatal("DISCORD_KEY or DISCORD_SECRET environment variables are not set")
	}

	goth.UseProviders(
		discord.New(discordKey, discordSecret, "http://localhost:8080/api/v1/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
	)
}
