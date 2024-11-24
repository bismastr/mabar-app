package handler

import (
	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/internal/user"
)

type Handler struct {
	bot           *bot.BotGamingSessionService
	gamingSession *gamingSession.GamingSessionService
	auth          *auth.AuthService
	user          *user.UserService
}

func NewHandler(bot *bot.BotGamingSessionService, gamingSession *gamingSession.GamingSessionService, auth *auth.AuthService, user *user.UserService) *Handler {
	return &Handler{
		bot:  bot,
		auth: auth,
		user: user,
	}
}
