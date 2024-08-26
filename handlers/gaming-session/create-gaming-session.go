package gaming_session

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionsList := i.ApplicationCommandData().Options
	var messageContent string

	if optionsList == nil {
		messageContent = "# Info Info Info Mabar dulu ga sih? @here"
	} else {
		messageContent = fmt.Sprintf("# Info Mabar %v dulu ga sih? @here", optionsList[0].Value)
	}

	if !mabarSession {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: messageContent,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🔥",
								},
								Label:    "Gas Join!",
								Style:    discordgo.PrimaryButton,
								CustomID: "mabar_yes",
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
		} else {
			mabarSession = true
		}
	} else {
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

}
