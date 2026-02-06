package message_components

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func NeedLoginMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	messageContent := "*Unable to join, user not yet registered*\n### Go register now to join mabar sessions! 🔥\n[Join Here](https://api-mabar.bismasatria.com/api/v1/auth/discord)\n"
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
