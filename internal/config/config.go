package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost              string
	Port                    string
	DiscordBotToken         string
	DiscordClientID         string
	DiscordClientSecret     string
	CookiesAuthSecret       string
	CookiesAuthAgeInSeconds int
	CookiesAuthIsSecure     bool
	CookiesAuthIsHttpOnly   bool
	SessionName             string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:              getEnv("PUBLIC_HOST", "http:localhost"),
		Port:                    getEnv("PORT", "8080"),
		CookiesAuthSecret:       getEnv("COOKIES_AUTH_SECRET", "some-very-secret-key"),
		CookiesAuthAgeInSeconds: getEnvAsInt("COOKIES_AUTH_AGE_IN_SECONDS", 86400*30), // 30 days
		CookiesAuthIsSecure:     getEnvAsBool("COOKIES_AUTH_IS_SECURE", false),
		CookiesAuthIsHttpOnly:   getEnvAsBool("COOKIES_AUTH_IS_HTTP_ONLY", false),
		DiscordClientID:         getEnvOrError("DISCORD_CLIENT_ID"),
		DiscordClientSecret:     getEnvOrError("DISCORD_CLIENT_SECRET"),
		DiscordBotToken:         fmt.Sprintf("Bot %v", getEnv("DISCORD_BOT_TOKEN", "token")),
		SessionName:             getEnv("SESSION_NAME", "session_user"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))

}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}

		return b
	}

	return fallback
}
