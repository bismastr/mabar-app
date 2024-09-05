package bot

import (
	"slices"

	"github.com/bismastr/discord-bot/components"
	"github.com/bismastr/discord-bot/internal/session"
	"github.com/bwmarrin/discordgo"
)

func GetOptionValueByName(i *discordgo.InteractionCreate, optionName string) any {
	var optionValue any

	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if option, ok := optionMap[optionName]; ok {
		optionValue = option.Value
	}

	return optionValue
}

func IsInSession(gs *session.GamingSession, userid string, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if slices.Contains(gs.MembersSession, userid) {
		components.AlreadyInSession(s, i)
		return true
	}

	return false
}
