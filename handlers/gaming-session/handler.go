package gaming_session

import (
	"context"

	"github.com/bismastr/discord-bot/db"
	"github.com/bwmarrin/discordgo"
)

var (
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context){
		"create-mabar": CreateGamingSession,
		"buyar-sek":    DeleteGamingSession,
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient){
		"mabar_yes": JoinGamingSession,
		"mabar_no":  DeclineGamingSession,
	}
)

func AddGamingSessionCommandData(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i, db, ctx)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i, db)
		}
	}
}
