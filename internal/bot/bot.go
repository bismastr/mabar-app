package bot

import (
	"fmt"

	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bwmarrin/discordgo"
)

type BotService struct {
	dg *discordgo.Session
}

func NewBotService(dg *discordgo.Session) *BotService {
	return &BotService{
		dg: dg,
	}
}

func (b *BotService) CreateGamingSession(gamingSession *gaming_session.GetGamingSessionResponse, channelId string) (*discordgo.Message, error) {
	content := fmt.Sprintf("# Info mabar? @here\n🎮 **Playing** 🎮\n%s \n\n🕐 On 🕐\n[Malam Ini]\n\n👥 Players 👥\n\n\n> MABAR ᴄʀᴇᴀᴛᴇᴅ ʙʏ <@%d>\n> Try mabar website: [Mabar Website](https://mabar.bism.app/)", gamingSession.Game.GameName.String, gamingSession.CreatedBy.DiscordUid.Int64)
	if gamingSession.Game.GameName.String == "" {
		content = fmt.Sprintf("# Info mabar? @here\n🎮 **Playing** 🎮\nBebas Asal Sopan \n\n🕐 On 🕐\n[Malam Ini]\n\n👥 Players 👥\n\n\n> MABAR ᴄʀᴇᴀᴛᴇᴅ ʙʏ <@%d>\n> Try mabar website: [Mabar Website](https://mabar.bism.app/)", gamingSession.CreatedBy.UserID.Int64)
	}

	customId := fmt.Sprintf("mabarv2_yes_%d", gamingSession.SessionID)

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
						CustomID: customId,
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

	res, err := b.dg.ChannelMessageSendComplex(channelId, message)
	if err != nil {
		return nil, err
	}

	return res, nil
}
