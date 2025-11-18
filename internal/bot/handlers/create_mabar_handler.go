package handlers

import (
	"context"
	"strconv"

	"github.com/bismastr/discord-bot/internal/bot/components/message_components"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type BotMessenger interface {
	CreateGamingSession(gamingSession *gaming_session.GetGamingSessionResponse, channelId string) (*discordgo.Message, error)
}

type CreateMabarHandler struct {
	userService *user.UserService
	sessionSvc  *gaming_session.GamingSessionService
	botService  BotMessenger
	ctx         context.Context
}

func NewCreateMabarHandler(
	userService *user.UserService,
	sessionSvc *gaming_session.GamingSessionService,
	botService BotMessenger,
	ctx context.Context,
) *CreateMabarHandler {
	return &CreateMabarHandler{
		userService: userService,
		sessionSvc:  sessionSvc,
		botService:  botService,
		ctx:         ctx,
	}
}

func (h *CreateMabarHandler) Name() string {
	return "create-mabar"
}

func (h *CreateMabarHandler) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)

	user, err := h.userService.GetUserByDiscordUID(h.ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			message_components.NeedLoginMessage(s, i)
		} else {
			message_components.ErrorMessage(s, i)
		}
		return err
	}

	createSession, err := h.sessionSvc.CreateGamingSession(h.ctx, &gaming_session.CreateGamingSessionRequest{
		IsFinish: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
		CreatedBy: user.ID,
		GameID:    i.ApplicationCommandData().Options[0].IntValue(),
	})
	if err != nil {
		message_components.ErrorMessage(s, i)
		return err
	}

	session, err := h.sessionSvc.GetGamingSessionById(h.ctx, createSession.ID)
	if err != nil {
		message_components.ErrorMessage(s, i)
		return err
	}

	_, err = h.botService.CreateGamingSession(session, i.ChannelID)
	if err != nil {
		message_components.ErrorMessage(s, i)
		return err
	}

	return nil
}
