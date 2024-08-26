package gaming_session

import (
	"context"

	"github.com/bismastr/discord-bot/db"
	"github.com/bwmarrin/discordgo"
)

func DeleteGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	if mabarSession {
		messageContent := "Buyar dulu kawanku, sampai jumpa di mabar berikutnya!👋"
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: messageContent,
			},
		})

		mabarSession = false
		membersSession = nil

		if err != nil {
			panic(err)
		}

	} else {
		messageContent := "❌ Tidak ada sesi mabar yang berjalan, kawan"
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: messageContent,
			},
		})

		if err != nil {
			panic(err)
		}
	}
}
