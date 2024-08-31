package components

import (
	"github.com/bwmarrin/discordgo"
)

func CreateSession(s *discordgo.Session, i *discordgo.InteractionCreate, id string, gameName string) {
	content := "## Ada info " + gameName + " hari ini? @here"
	if gameName == "" {
		content = "## Ada info Minecraft or CS nanti malam? @here"
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
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
		},
	})
	if err != nil {
		panic(err)
	}
}

func UnableCreateSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "❌ Sesi mabar sudah ada, kawanku",
		},
	})

	if err != nil {
		panic(err)
	}
}
