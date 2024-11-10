package bot

import (
	"slices"
	"strings"

	"github.com/bismastr/discord-bot/internal/bot/components"
	"github.com/bismastr/discord-bot/internal/gamingSession"
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

func IsInSession(gs *gamingSession.GamingSession, userid string, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if slices.Contains(gs.MembersSession, userid) {
		components.AlreadyInSession(s, i)
		return true
	}

	return false
}

func getPrefix(i *discordgo.InteractionCreate) string {
	customID := i.MessageComponentData().CustomID
	split := strings.Split(customID, "_")
	prefix := split[0] + "_" + split[1]

	return prefix
}
