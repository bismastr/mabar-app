package bot

import (
	"context"

	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RegisterHandler() {
	//Action Handler
	b.Dg.AddHandler(b.interactionHandler)
}

func (b *Bot) interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	repository := gamingSession.NewRepositoryImpl(b.database)
	h := NewActionHandlerCtrl(gamingSession.NewGamingSessionService(repository), context.Background())
	var (
		commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"create-mabar": h.CreateSession,
			// "list-mabar":   ListSession,
		}
		componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"mabar_yes":  h.JoinGamingSession,
			"mabar_no":   h.DeclineGamingSession,
			"init_mabar": h.InitMabar,
		}
	)

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:
		prefix := getPrefix(i)

		if h, ok := componentsHandlers[prefix]; ok {
			h(s, i)
		}
	}

}
