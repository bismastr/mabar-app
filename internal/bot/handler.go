package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RegisterHandler(h *ActionHandlerCtrl) {
	b.Dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		b.interactionHandler(h, s, i)
	})
}

func (b *Bot) interactionHandler(h *ActionHandlerCtrl, s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"create-mabar":          h.CreateMabar,
			"ask-ai":                h.GenerateContent,
			"create-daily-cs-alert": h.CreateSchedulerCsItems,
		}
		componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"mabar_no":    h.DeclineGamingSession,
			"mabarv2_yes": h.JoinGamingSessionV2,
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
	case discordgo.InteractionApplicationCommandAutocomplete:
		h.CsItemsAutocomplete(s, i)
	}

}
