package gaming_session

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	membersSession []string
	mabarSession   bool
)

func GenerateMemberMention(members []string) string {
	result := ""
	for _, s := range members {
		result += fmt.Sprintf("<@%v>", s)
	}
	return result
}

func CheckJoin(userId string) bool {
	for _, u := range membersSession {
		if u == userId {
			return true
		}
	}
	return false
}

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
