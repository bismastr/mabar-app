package message_components

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendMessage(s *discordgo.Session, i *discordgo.InteractionCreate, messageContent string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: messageContent,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
