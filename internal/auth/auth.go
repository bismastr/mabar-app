package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

func NewAuth() {
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET environment variable is not set")
	}

	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.MaxAge(86400 * 30)

	store.Options.HttpOnly = true
	store.Options.Path = "/"
	store.Options.Secure = false
	store.Options.SameSite = http.SameSiteLaxMode

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
