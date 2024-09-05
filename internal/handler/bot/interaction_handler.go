package bot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bismastr/discord-bot/components"
	"github.com/bismastr/discord-bot/internal/session"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
)

type ActionHandlerCtrl struct {
	gamingSessionService *session.GamingSessionService
	ctx                  context.Context
}

func NewActionHandlerCtrl(gamingSessionService *session.GamingSessionService, ctx context.Context) *ActionHandlerCtrl {
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
	updateMember := session.GamingSession{
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
	gameName := ""
	if optionValue := GetOptionValueByName(i, "nama-permainan"); optionValue != nil {
		gameName = optionValue.(string)
	}

	session := session.GamingSession{
		CreatedAt: time.Now().String(),
		CreatedBy: &session.CreatedBy{
			Id:       i.Member.User.ID,
			Username: i.Member.Nick,
		},
		SessionEnd:   "", //Need to add session
		SessionStart: "",
		GameName:     gameName,
		IsFinish:     false,
	}

	id, err := a.gamingSessionService.CreateGamingSession(a.ctx, session)
	if err != nil {
		panic(err)
	}

	components.CreateSession(s, i, id, gameName)
}
