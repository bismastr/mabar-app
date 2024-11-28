package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bismastr/discord-bot/internal/bot/components"
	"github.com/bismastr/discord-bot/internal/bot/components/message_components"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
)

type ActionHandlerCtrl struct {
	gamingSessionService *gamingSession.GamingSessionService
	userService          *user.UserService
	gamingSession        *gaming_session.GamingSessionService
	ctx                  context.Context
}

func NewActionHandlerCtrl(
	gamingSessionService *gamingSession.GamingSessionService,
	userService *user.UserService,
	gamingSession *gaming_session.GamingSessionService,
	ctx context.Context) *ActionHandlerCtrl {
	return &ActionHandlerCtrl{
		gamingSessionService: gamingSessionService,
		userService:          userService,
		gamingSession:        gamingSession,
		ctx:                  ctx,
	}
}

func (a *ActionHandlerCtrl) JoinGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userid := i.Member.User.ID
	customId := i.MessageComponentData().CustomID
	split := strings.Split(customId, "_")
	refId := split[2]

	currentRef, err := a.gamingSessionService.GetGamingSessionByRefId(a.ctx, refId)
	if err != nil {
		panic(err)
	}

	if IsInSession(currentRef, userid, s, i) {
		return
	}

	updateMember := gamingSession.GamingSession{
		MembersSession: append(currentRef.MembersSession, userid),
	}
	err = a.gamingSessionService.UpdateGamingSessionByRefId(a.ctx, refId, updateMember)
	if err != nil {
		panic(err)
	}

	components.JoinSession(s, i, userid, GenerateMemberMention(updateMember.MembersSession))
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

func (a *ActionHandlerCtrl) CreateSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	gameName := []discordgo.PollAnswer{}
	session := gamingSession.GamingSession{
		CreatedAt: time.Now().String(),
		CreatedBy: &gamingSession.CreatedBy{
			Id:       i.Member.User.ID,
			Username: i.Member.Nick,
		},
		SessionEnd:   "", //Need to add session
		SessionStart: "",
		IsFinish:     false,
	}

	for _, v := range i.ApplicationCommandData().Options {
		if v.StringValue() != "" {
			gameName = append(gameName, discordgo.PollAnswer{
				Media: &discordgo.PollMedia{
					Text: v.StringValue(),
				},
			})
		}
	}

	var gameText string
	switch len(gameName) {
	case 1:
		gameText = gameName[0].Media.Text
		session.GameName = gameText
		id, err := a.gamingSessionService.CreateGamingSession(a.ctx, session)
		if err != nil {
			panic(err)
		}
		components.CreateSession(s, i, id, gameText)
	case 0:
		session.GameName = ""
		id, err := a.gamingSessionService.CreateGamingSession(a.ctx, session)
		if err != nil {
			panic(err)
		}
		components.CreateSession(s, i, id, "")
	default:
		id, err := a.gamingSessionService.CreateGamingSession(a.ctx, session)
		if err != nil {
			panic(err)
		}
		components.CreateSessionPoll(s, i, gameName, id)
	}
}

func (a *ActionHandlerCtrl) InitMabar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	customId := i.MessageComponentData().CustomID
	split := strings.Split(customId, "_")
	refId := split[2]

	currentRef, err := a.gamingSessionService.GetGamingSessionByRefId(a.ctx, refId)
	if err != nil {
		panic(err)
	}

	if currentRef.CreatedBy.Id != i.Member.User.ID {
		components.UnableCreateSession(s, i)
		return
	}

	m, _ := s.PollExpire(i.ChannelID, i.Message.ID)

	var userWinning []*discordgo.User
	var gameName string
	for _, v := range m.Poll.Answers {
		user, _ := s.PollAnswerVoters(i.ChannelID, i.Message.ID, v.AnswerID)

		if len(user) > len(userWinning) {
			userWinning = user
			gameName = v.Media.Text
		}
	}

	updateGamingSession := gamingSession.GamingSession{
		GameName: gameName,
	}

	for _, v := range userWinning {
		updateGamingSession.MembersSession = append(updateGamingSession.MembersSession, v.ID)
	}

	err = a.gamingSessionService.UpdateGamingSessionByRefId(a.ctx, refId, updateGamingSession)
	if err != nil {
		panic(err)
	}

	components.InitMabar(s, i, gameName, GenerateMemberMention(updateGamingSession.MembersSession))
	defer s.ChannelMessageDelete(i.ChannelID, i.Message.ID)
}
