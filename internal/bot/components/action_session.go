package components

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func JoinSession(s *discordgo.Session, i *discordgo.InteractionCreate, userId string, memberMentioned string) {
	messageContent := fmt.Sprintf("<@%v> join abang quh 🥳\n\nArek-arek sing join 👥:%v  ", userId, memberMentioned)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: messageContent,
		},
	})
	if err != nil {
		panic(err)
	}
}

func AlreadyInSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Kamu udah join sesi, abangkuh",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		panic(err)
	}
}

func InitMabar(s *discordgo.Session, i *discordgo.InteractionCreate, gameName string, members string) {
	content := fmt.Sprintf("# Mabar Started! 🔥 \n## Playing %v\n\nAyo join bang! %v", gameName, members)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		panic(err)
	}
}
