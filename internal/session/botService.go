package session

import "github.com/bwmarrin/discordgo"

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

func (b *BotGamingSessionService) CreateGamingSession(id string, gameName string) (*discordgo.Message, error) {
	b.dg.ChannelMessageSend("1276782792876888075", "test create gaming session")

	content := "## Ada info " + gameName + " hari ini? @here"
	if gameName == "" {
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
