package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bismastr/discord-bot/internal/bot/components/message_components"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/llm"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ActionHandlerCtrl struct {
	userService   *user.UserService
	gamingSession *gaming_session.GamingSessionService
	BotService    *BotService
	llmService    *llm.LlmService
	ctx           context.Context
}

func NewActionHandlerCtrl(
	userService *user.UserService,
	gamingSession *gaming_session.GamingSessionService,
	botService *BotService,
	llmService *llm.LlmService,
	ctx context.Context) *ActionHandlerCtrl {
	return &ActionHandlerCtrl{
		userService:   userService,
		gamingSession: gamingSession,
		BotService:    botService,
		ctx:           ctx,
		llmService:    llmService,
	}
}

func (a *ActionHandlerCtrl) GenerateContent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp, err := a.llmService.GetGenerateResponse(a.ctx, "7 tambah 8 berapa?")
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	message_components.SendMessage(s, i, fmt.Sprintf("AI Response: %v", resp))
}
func (a *ActionHandlerCtrl) JoinGamingSessionV2(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)
	customId := i.MessageComponentData().CustomID
	split := strings.Split(customId, "_")
	id, _ := strconv.ParseInt(split[2], 10, 64)

	user, err := a.userService.GetUserByDiscordUID(a.ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			message_components.NeedLoginMessage(s, i)
		} else {
			message_components.ErrorMessage(s, i)
		}
		return
	}

	err = a.gamingSession.InsertUserJoinSession(a.ctx, user.ID, id)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	response, err := a.gamingSession.GetGamingSessionById(a.ctx, id)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	message_components.JoinSessionV2(s, i, userId, response)
}

func (a *ActionHandlerCtrl) DeclineGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userid := i.Member.User.ID
	noJoin := fmt.Sprintf("<@%v> tidak join duls, kecewaaaa sangat berat!", userid)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: noJoin,
		},
	})
	if err != nil {
		panic(err)
	}
}

func (a *ActionHandlerCtrl) CreateMabar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)

	user, err := a.userService.GetUserByDiscordUID(a.ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			message_components.NeedLoginMessage(s, i)
		} else {
			message_components.ErrorMessage(s, i)
		}
		return
	}

	createSession, err := a.gamingSession.CreateGamingSession(a.ctx, &gaming_session.CreateGamingSessionRequest{
		IsFinish: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
		CreatedBy: user.ID,
		GameID:    i.ApplicationCommandData().Options[0].IntValue(),
	})
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	session, err := a.gamingSession.GetGamingSessionById(a.ctx, createSession.ID)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	_, err = a.BotService.CreateGamingSession(session, i.ChannelID)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}
}
