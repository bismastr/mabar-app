package gaming_session

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"create-mabar": CreateGamingSession,
		"buyar-sek":    DeleteGamingSession,
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"mabar_yes": JoinGamingSession,
		"mabar_no":  DeclineGamingSession,
	}
)

func AddGamingSessionCommandData(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}
	}
}
