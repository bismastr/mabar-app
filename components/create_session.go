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

func CreateSessionPoll(s *discordgo.Session, i *discordgo.InteractionCreate, answers []discordgo.PollAnswer, id string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "## Ada info permainan hari ini? @here",
			Poll: &discordgo.Poll{
				Question: discordgo.PollMedia{
					Text: "Main apa hari ini?",
				},
				Answers:          answers,
				AllowMultiselect: true,
				LayoutType:       discordgo.PollLayoutTypeDefault,
				Duration:         1,
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "🔥",
							},
							Label:    "Mulai Sesi Mabar",
							Style:    discordgo.PrimaryButton,
							CustomID: "init_mabar_" + id,
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
			Content: "❌ You cannot start other people's session. Only session creator can start the session.",
		},
	})

	if err != nil {
		panic(err)
	}
}
