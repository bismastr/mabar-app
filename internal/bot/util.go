package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func getPrefix(i *discordgo.InteractionCreate) string {
	customID := i.MessageComponentData().CustomID
	split := strings.Split(customID, "_")
	prefix := split[0] + "_" + split[1]

	return prefix
}
