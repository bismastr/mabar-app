package message_components

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ErrorMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	messageContent := "Something Went Wrong, try again later"
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: messageContent,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
