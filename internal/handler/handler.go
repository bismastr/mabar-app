package handler

import (
	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/user"
)

type Handler struct {
	bot           *bot.BotGamingSessionService
	gamingSession *gamingSession.GamingSessionService
	auth          *auth.AuthService
	user          *user.UserService
	gaming        *gaming_session.GamingSessionService
}

func NewHandler(bot *bot.BotGamingSessionService, gamingSession *gamingSession.GamingSessionService, auth *auth.AuthService, user *user.UserService, gaming *gaming_session.GamingSessionService) *Handler {
	return &Handler{
		bot:           bot,
		gamingSession: gamingSession,
		auth:          auth,
		user:          user,
		gaming:        gaming,
	}
}
