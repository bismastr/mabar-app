package gaming_session

import (
	"context"
	"strings"

	"github.com/bismastr/discord-bot/db"
	"github.com/bwmarrin/discordgo"
)

var (
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context){
		"create-mabar": CreateSession,
		"list-mabar":   ListSession,
		// "buyar-sek":    DeleteGamingSession,
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context){
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
		customID := i.MessageComponentData().CustomID
		split := strings.Split(customID, "_")
		prefix := split[0] + "_" + split[1]

		if h, ok := componentsHandlers[prefix]; ok {
			h(s, i, db, ctx)
		}
	}
}
