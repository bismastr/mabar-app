package handler

import (
	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/gamingSession"
)

type Handler struct {
	bot           *bot.BotGamingSessionService
	gamingSession *gamingSession.GamingSessionService
	auth          *auth.AuthService
}

func NewHandler(bot *bot.BotGamingSessionService, gamingSession *gamingSession.GamingSessionService, auth *auth.AuthService) *Handler {
	return &Handler{
		bot:           bot,
		gamingSession: gamingSession,
		auth:          auth,
	}
}
