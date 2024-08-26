package components

import "github.com/bwmarrin/discordgo"

func CreateSession(s *discordgo.Session, i *discordgo.InteractionCreate, id string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "# Info Info Info Mabar dulu ga sih? @here",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "🔥",
							},
							Label:    "Gas Join!",
							Style:    discordgo.PrimaryButton,
							CustomID: "mabar_yes_" + id,
						},
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "❌",
							},
							Label:    "Skip duls",
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
