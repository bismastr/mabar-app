package handlers

import (
	"context"
	"strconv"
	"strings"

	"github.com/bismastr/discord-bot/internal/bot/components/message_components"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
)

type JoinSessionHandler struct {
	userService *user.UserService
	sessionSvc  *gaming_session.GamingSessionService
	ctx         context.Context
}

func NewJoinSessionHandler(
	userService *user.UserService,
	sessionSvc *gaming_session.GamingSessionService,
	ctx context.Context,
) *JoinSessionHandler {
	return &JoinSessionHandler{
		userService: userService,
		sessionSvc:  sessionSvc,
		ctx:         ctx,
	}
}

func (h *JoinSessionHandler) CustomIDPrefix() string {
	return "mabarv2_yes"
}

func (h *JoinSessionHandler) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)

	customId := i.MessageComponentData().CustomID
	split := strings.Split(customId, "_")
	sessionId, _ := strconv.ParseInt(split[2], 10, 64)

	user, err := h.userService.GetUserByDiscordUID(h.ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			message_components.NeedLoginMessage(s, i)
		} else {
			message_components.ErrorMessage(s, i)
		}
		return err
	}

	err = h.sessionSvc.InsertUserJoinSession(h.ctx, user.ID, sessionId)
	if err != nil {
		message_components.ErrorMessage(s, i)
		return err
	}

	response, err := h.sessionSvc.GetGamingSessionById(h.ctx, sessionId)
	if err != nil {
		message_components.ErrorMessage(s, i)
		return err
	}

	message_components.JoinSessionV2(s, i, userId, response)
	return nil
}
