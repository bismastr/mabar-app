package gaming_session

import (
	"context"

	"github.com/bismastr/discord-bot/db"
	"github.com/bismastr/discord-bot/model"
	"github.com/bwmarrin/discordgo"
)

func CreateGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	session := model.GamingSession{
		CreatedAt: "bisma testing create session",
	}

	db.CreateGamingSession(ctx, session)

	if !mabarSession {
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
