package gaming_session

import (
	"context"

	"github.com/bismastr/discord-bot/components"
	"github.com/bismastr/discord-bot/db"
	"github.com/bwmarrin/discordgo"
)

func ListSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	result, err := db.ReadGamingSession(ctx)
	if err != nil {
		panic(err)
	}

	components.ListSession(s, i, result)
}
