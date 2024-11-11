package auth

import (
	"log"
	"os"

	"github.com/bismastr/discord-bot/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

type AuthService struct{}

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
