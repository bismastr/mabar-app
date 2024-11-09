package session

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type BotGamingSessionService struct {
	repository FirestoreRepositorySession
	dg         *discordgo.Session
}

func NewBotGamingSessionService(repo FirestoreRepositorySession, dg *discordgo.Session) *BotGamingSessionService {
	return &BotGamingSessionService{
		repository: repo,
		dg:         dg,
	}
}

func (b *BotGamingSessionService) CreateGamingSession(id string, gamingSession *GamingSession) (*discordgo.Message, error) {
	b.dg.ChannelMessageSend("1276782792876888075", "test create gaming session")

	content := fmt.Sprintf("## Ada info %s hari ini? @here \n created by @%s", gamingSession.GameName, gamingSession.CreatedBy.Username)
	if gamingSession.GameName == "" {
		content = "## Ada info permainan hari ini? @here"
	}
	message := &discordgo.MessageSend{
		Content: content,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "🔥",
						},
						Label:    "Gas!",
						Style:    discordgo.PrimaryButton,
						CustomID: "mabar_yes_" + id,
					},
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "❌",
						},
						Label:    "Skip",
						Style:    discordgo.SecondaryButton,
						CustomID: "mabar_no",
					},
				},
			},
		},
	}

	res, err := b.dg.ChannelMessageSendComplex("1276782792876888075", message)
	if err != nil {
		return nil, err
	}

	return res, nil
}
