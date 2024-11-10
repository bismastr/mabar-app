package bot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bismastr/discord-bot/internal/bot/components"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
)

type ActionHandlerCtrl struct {
	gamingSessionService *gamingSession.GamingSessionService
	ctx                  context.Context
}

func NewActionHandlerCtrl(gamingSessionService *gamingSession.GamingSessionService, ctx context.Context) *ActionHandlerCtrl {
	return &ActionHandlerCtrl{
		gamingSessionService: gamingSessionService,
		ctx:                  ctx,
	}
}

// JoinGamingSession is a function to handle user who join the gaming session.
// This function will check is user already in session or not. If not, then it will update the gaming session.
// After that, it will respond to user with a message that user join the gaming session.
func (a *ActionHandlerCtrl) JoinGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//GetUserId and Refid
	userid := i.Member.User.ID
	customId := i.MessageComponentData().CustomID
	split := strings.Split(customId, "_")
	refId := split[2]
	//GetCurrentGamingSession
	currentRef, err := a.gamingSessionService.GetGamingSessionByRefId(a.ctx, refId)
	if err != nil {
		panic(err)
	}

	if IsInSession(currentRef, userid, s, i) {
		return
	}

	//UpdateMember
	updateMember := gamingSession.GamingSession{
		MembersSession: append(currentRef.MembersSession, userid),
	}
	err = a.gamingSessionService.UpdateGamingSessionByRefId(a.ctx, refId, updateMember)
	if err != nil {
		panic(err)
	}

	components.JoinSession(s, i, userid, utils.GenerateMemberMention(updateMember.MembersSession))
}

// DeclineGamingSession is a function to handle user who decline the gaming session.
// This function will respond to user with a message that user decline the gaming session.
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

// CreateSession is a function to handle user who create the gaming session.
// This function will check is user already in session or not. If not, then it will create the gaming session.
// After that, it will respond to user with a message that user create the gaming session.
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
	//

	//Switch
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
	//RefId
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

	components.InitMabar(s, i, gameName, utils.GenerateMemberMention(updateGamingSession.MembersSession))
	defer s.ChannelMessageDelete(i.ChannelID, i.Message.ID)
}
