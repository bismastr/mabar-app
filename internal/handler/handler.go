package handler

import (
	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/user"
)

type Handler struct {
	bot            *bot.BotService
	auth           *auth.AuthService
	user           *user.UserService
	gaming_session *gaming_session.GamingSessionService
}

func NewHandler(
	bot *bot.BotService,
	auth *auth.AuthService,
	user *user.UserService,
	gaming_session *gaming_session.GamingSessionService) *Handler {
	return &Handler{
		bot:            bot,
		auth:           auth,
		user:           user,
		gaming_session: gaming_session,
	}
}
